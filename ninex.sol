pragma solidity ^0.4.0;

contract Ninex {
  // Owner of the contract.
  address public mOwner;

  // Commitment to the value being guessed.  It's the sha256 hash of a
  // bytes value consisting of 78 ASCII base 10 digits. Players guess
  // the final N digits to win (their stake) * 9 * 10^(N-1). (So, 9x
  // if N is 1, which is where the name of the contract comes from.)
  bytes32 public mCommitment;

  // Preimage is the string that produces mCommitment. It's set only
  // after reveal is called and is reset when setCommitment() is
  // called.
  bytes public mPreimage;

  // Amount available for covering new guesses and for withdrawals.
  // Invariant: bank+escrow+payments == contract balance
  uint public mBank;

  struct Guess {
    address guesser;
    bytes   digits;
    uint    escrow;
  }

  Guess[] public mGuesses;

  // States:
  //  0. all times zero [next: 1]
  //  1. mCommitmentSetTime > 0 [next: 2 (guess)]
  //  2. mCommitmentSetTime > 0 && mFirstGuessTime > 0 && mLastGuessTime > 0 [next: 0 (timeout) or 3 (reveal)]
  //  3a. all times > 0 [next: 1 (setCommitment)]
  uint public mCommitmentSetTime;
  uint public mFirstGuessTime;
  uint public mLastGuessTime;
  uint public mRevealedTime;

  // Params
  uint public mAfterCommitmentDelaySecs; // No guess too soon after a new commitment.
  uint public mGuessWindowSecs;          // After first guess, how long to accept additional guesses.
  uint public mAfterLastGuessDelaySecs;  // No reveal too soon after the last guess.
  uint public mAfterRevealDelaySecs;     // No new commitment too soon after a reveal.
  uint public mRevealTimeoutSecs;        // Right or wrong, guessers win if no reveal in this interval after first guess.
  uint public mMinGuessWei;              // Smallest amount (in wei) to guess with.

  // Payments for winning guesses accrue here.
  mapping (address => uint) public mPayments;

  // Constructor.
  function Ninex() {
    mOwner = msg.sender;

    mAfterCommitmentDelaySecs = 600;
    mGuessWindowSecs          = 900;
    mAfterLastGuessDelaySecs  = 600;
    mAfterRevealDelaySecs     = 600;
    mRevealTimeoutSecs        = 86400;
    mMinGuessWei              = 1000000;

    mBank              = 0;

    mCommitmentSetTime = 0;
    mFirstGuessTime    = 0;
    mLastGuessTime     = 0;
    mRevealedTime      = 0;
  }

  // Anyone may add funds to the bank at any time.
  function fund() external payable {
    mBank += msg.value;
  }

  // The owner may withdraw funds that aren't locked in escrow or
  // unclaimed payments.
  function withdraw(uint amount) external {
    require (msg.sender == mOwner);
    require (amount <= mBank);
    mBank -= amount;
    mOwner.transfer(amount);
  }

  event evGuess(address guesser, bytes32 commitment, bytes digits, uint stake);

  function guess(bytes32 commitment, bytes digits) external payable {
    // A commitment must be set.
    require (mCommitmentSetTime > 0);

    // No preimage must be revealed.
    require (mRevealedTime == 0);

    // Must not guess too soon after a new commitment is set.
    require (now >= (mCommitmentSetTime + mAfterCommitmentDelaySecs));

    // Must guess within the guessing window.
    require ((mFirstGuessTime == 0) || (now < (mFirstGuessTime + mGuessWindowSecs)));

    require (digits.length >= 1);
    require (digits.length <= 78);
    require (msg.value >= mMinGuessWei);

    // Must not break the bank.
    var potentialPayment = msg.value * 9 * (10 ** (digits.length - 1)) - msg.value;
    require (potentialPayment <= mBank);

    // Must be guessing against the expected commitment.
    var sameCommitment = true;
    for (uint i = 0; i < 32; i++) {
      if (commitment[i] != mCommitment[i]) {
        sameCommitment = false;
        break;
      }
    }
    require (sameCommitment);

    mGuesses.push(Guess({
      guesser: msg.sender,
      digits: digits,
      escrow: msg.value + potentialPayment
    }));

    mBank -= potentialPayment;

    if (mFirstGuessTime == 0) {
      mFirstGuessTime = now;
    }
    mLastGuessTime = now;

    evGuess(msg.sender, commitment, digits, msg.value);
  }

  event evNewCommitment(bytes32 commitment);

  function setCommitment(bytes32 commitment) external {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    require ((mCommitmentSetTime == 0) || (mRevealedTime > 0));
    require ((mRevealedTime == 0) || (now >= mRevealedTime + mAfterRevealDelaySecs));

    mCommitment = commitment;
    mPreimage = '';

    mCommitmentSetTime = now;

    mRevealedTime = 0;
    mFirstGuessTime = 0;
    mLastGuessTime = 0;

    evNewCommitment(commitment);
  }

  event evWin(address guesser, uint value, bytes digits, bytes32 commitment);
  event evLose(address guesser, bytes digits, bytes32 commitment);
  event evRevealed(bytes32 commitment, bytes preimage);

  function reveal(bytes preimage) external {
    require (msg.sender == mOwner);
    require (mCommitmentSetTime > 0);
    require (mRevealedTime == 0);

    // Preimage must be 78 bytes long (each byte is a base-10 digit
    // for 3.3 bits of entropy; 78 bytes gives > 256 bits of entropy).
    require (preimage.length == 78);

    // Don't reveal too soon after the last guess.
    require ((mLastGuessTime == 0) || (now >= (mLastGuessTime + mAfterLastGuessDelaySecs)));

    // Don't reveal after the timeout.
    require ((mFirstGuessTime == 0) || (now < (mFirstGuessTime + mRevealTimeoutSecs)));

    // Must reveal a valid preimage.
    require (sha256(preimage) == mCommitment);

    for (uint i = 0; i < mGuesses.length; i++) {
      var digitsMatch = true;
      var offset = 78 - mGuesses[i].digits.length;
      for (uint j = offset; j < 78; j++) {
        if (mGuesses[i].digits[j-offset] != preimage[j]) {
          digitsMatch = false;
          break;
        }
      }
      if (digitsMatch) {
        mPayments[mGuesses[i].guesser] += mGuesses[i].escrow;
        evWin(mGuesses[i].guesser, mGuesses[i].escrow, mGuesses[i].digits, mCommitment);
      } else {
        mBank += mGuesses[i].escrow;
        evLose(mGuesses[i].guesser, mGuesses[i].digits, mCommitment);
      }
    }
    mGuesses.length = 0;

    mRevealedTime = now;
    mPreimage = preimage;

    evRevealed(mCommitment, preimage);
  }

  event evWinByDefault(address guesser, uint value);

  // Anyone may move the escrowed amount to a payment after the reveal
  // timeout elapses.
  function timeout() external {
    require (mCommitmentSetTime > 0);
    require (mFirstGuessTime > 0);
    require (mRevealedTime == 0);
    require (now >= (mFirstGuessTime + mRevealTimeoutSecs));

    for (uint i = 0; i < mGuesses.length; i++) {
      mPayments[mGuesses[i].guesser] += mGuesses[i].escrow;
      evWinByDefault(mGuesses[i].guesser, mGuesses[i].escrow);
    }
    mGuesses.length = 0;

    mCommitmentSetTime = 0;
    mFirstGuessTime = 0;
    mLastGuessTime = 0;
    mRevealedTime = 0;
  }

  function collect() external {
    var amount = mPayments[msg.sender];
    if (amount > 0) {
      mPayments[msg.sender] = 0;
      msg.sender.transfer(amount);
    }
  }

  function setAfterCommitmentDelay(uint val) {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    mAfterCommitmentDelaySecs = val;
  }

  function setGuessWindow(uint val) {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    mGuessWindowSecs = val;
  }

  function setAfterLastGuessDelay(uint val) {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    mAfterLastGuessDelaySecs = val;
  }

  function setAfterRevealDelay(uint val) {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    mAfterRevealDelaySecs = val;
  }

  function setRevealTimeout(uint val) {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    mRevealTimeoutSecs = val;
  }

  function setMinGuessWei(uint val) {
    require (msg.sender == mOwner);
    require (mGuesses.length == 0);
    mMinGuessWei = val;
  }

  function numGuesses() view returns (uint) {
    return mGuesses.length;
  }
}
