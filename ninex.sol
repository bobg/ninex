pragma solidity ^0.4.0;

contract Ninex {
  // Owner of the contract.
  address public mOwner;

  // Commitment to the value being guessed.  It's the sha256 hash of a
  // bytes value consisting of 78 ASCII base 10 digits. Players guess
  // the final N digits to win (their stake) * 9 * 10^(N-1). (So, 9x
  // if N is 1, which is where the name of the contract comes from.)
  bytes32 public mCommitment;

  // Amount available for covering new guesses and for withdrawals.
  // Invariant: bank+escrow+payments == contract balance
  uint public mBank;

  // Amount locked up (potential payment) after a guess is made and
  // before the outcome is revealed.
  // Invariant: mEscrow>0 iff there is a pending guess.
  uint public mEscrow;

  // When a guess is pending:
  address public mGuessedBy;   // The guesser.
  bytes public mGuessedDigits; // The digits of the guess.

  // At most one is nonzero.
  uint public mCommitmentSetTime;
  uint public mGuessedTime;
  uint public mRevealedTime;

  // Params
  uint public mAfterCommitmentDelaySecs; // No guess too soon after a new commitment.
  uint public mAfterGuessDelaySecs;      // No reveal too soon after a guess.
  uint public mAfterRevealDelaySecs;     // No new commitment too soon after a reveal.
  uint public mRevealTimeoutSecs;        // Right or wrong, guesser wins if no reveal in this interval.
  uint public mMinGuess;                 // Smallest amount (in wei) to guess with.

  // Payments for winning guesses accrue here.
  mapping (address => uint) public mPayments;

  // Constructor.
  function Ninex() {
    mOwner = msg.sender;

    mAfterCommitmentDelaySecs = 600;
    mAfterGuessDelaySecs      = 600;
    mAfterRevealDelaySecs     = 600;
    mRevealTimeoutSecs        = 86400;
    mMinGuess                 = 1000000;

    mBank              = 0;
    mEscrow            = 0;
    mGuessedTime       = 0;
    mRevealedTime      = 0;
    mCommitmentSetTime = 0;
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

  function guess(bytes32 commitment, bytes digits) external payable {
    // There must be no pending guess.
    require (mEscrow == 0);

    // A commitment must be set.
    require (mCommitmentSetTime > 0);

    require (digits.length >= 1);
    require (digits.length <= 78);
    require (msg.value >= mMinGuess);

    // Must not break the bank.
    var potentialPayment = msg.value * 9 * (10 ** (digits.length - 1)) - msg.value;
    require (potentialPayment <= mBank);

    // Must not guess too soon after a new commitment is set.
    require (now >= mCommitmentSetTime + mAfterCommitmentDelaySecs);

    // Must be guessing against the expected commitment.
    var sameCommitment = true;
    for (uint i = 0; i < 32; i++) {
      if (commitment[i] != mCommitment[i]) {
        sameCommitment = false;
        break;
      }
    }
    require (sameCommitment);

    mGuessedBy = msg.sender;
    mGuessedDigits = digits;

    mBank -= potentialPayment;
    mEscrow += potentialPayment;
    mEscrow += msg.value;

    mGuessedTime = now;
    mCommitmentSetTime = 0;
    mRevealedTime = 0;
  }

  function setCommitment(bytes32 commitment) external {
    require (msg.sender == mOwner);
    require (mEscrow == 0);
    require (mCommitmentSetTime == 0);
    require (now >= mRevealedTime + mAfterRevealDelaySecs);

    mCommitment = commitment;

    mCommitmentSetTime = now;
    mRevealedTime = 0;
    mGuessedTime = 0;
  }

  function reveal(bytes preimage) external {
    require (msg.sender == mOwner);

    // Preimage must be 78 bytes long (each byte is a base-10 digit
    // for 3.3 bits of entropy; 78 bytes gives > 256 bits of entropy).
    require (preimage.length == 78);

    require (sha256(preimage) == mCommitment);

    if (mEscrow > 0) {
      // A guess is pending.

      require (now >= (mGuessedTime + mAfterGuessDelaySecs));
      require (now < (mGuessedTime + mRevealTimeoutSecs));

      var digitsMatch = true;
      var offset = 78 - mGuessedDigits.length;
      for (uint i = offset; i < 78; i++) {
        if (mGuessedDigits[i-offset] != preimage[i]) {
          digitsMatch = false;
          break;
        }
      }
      if (digitsMatch) {
        mPayments[mGuessedBy] += mEscrow;
      } else {
        mBank += mEscrow;
      }
      mEscrow = 0;
    }

    mRevealedTime = now;
    mCommitmentSetTime = 0;
    mGuessedTime = 0;
  }

  // Anyone may move the escrowed amount to a payment after the reveal
  // timeout elapses.
  function timeout() external {
    require (mEscrow > 0);
    require (now >= (mGuessedTime + mRevealTimeoutSecs));

    mPayments[mGuessedBy] += mEscrow;
    mEscrow = 0;

    mGuessedTime = 0;
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
    require (mEscrow == 0);
    mAfterCommitmentDelaySecs = val;
  }

  function setAfterGuessDelay(uint val) {
    require (msg.sender == mOwner);
    require (mEscrow == 0);
    mAfterGuessDelaySecs = val;
  }

  function setAfterRevealDelay(uint val) {
    require (msg.sender == mOwner);
    require (mEscrow == 0);
    mAfterRevealDelaySecs = val;
  }

  function setRevealTimeout(uint val) {
    require (msg.sender == mOwner);
    require (mEscrow == 0);
    mRevealTimeoutSecs = val;
  }

  function setMinGuess(uint val) {
    require (msg.sender == mOwner);
    require (mEscrow == 0);
    mMinGuess = val;
  }
}
