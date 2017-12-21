var Ninex = {
  init: function(infoURL) {
    Ninex.infoURL = infoURL;

    Ninex.initializingEl = document.getElementById('initializing');
    Ninex.nocommitmentEl = document.getElementById('nocommitment');
    Ninex.okguessEl = document.getElementById('okguess');
    Ninex.noguessEl = document.getElementById('noguess');
    Ninex.staleEl = document.getElementById('stale');
    Ninex.commitmentEl = document.getElementById('commitment');
    Ninex.digitsEl = document.getElementById('digits');
    Ninex.stakeEl = document.getElementById('stake');
    Ninex.maxstakeEl = document.getElementById('maxstake');
    Ninex.payoutEl = document.getElementById('payout');
    Ninex.stakeinfoEl = document.getElementById('stakeinfo');

    var update = function() {
      var now = Math.floor(Date.now()/1000); // seconds since unix epoch
      var show = function(name) {
        Ninex.initializingEl.style.display = 'none';
        Ninex.nocommitmentEl.style.display = (name == 'nocommitment') ? 'block' : 'none';
        Ninex.okguessEl.style.display = (name == 'okguess') ? 'block' : 'none';
        Ninex.noguessEl.style.display = (name == 'noguess') ? 'block' : 'none';
        Ninex.staleEl.style.display = (name == 'stale') ? 'block' : 'none';
      };
      Ninex.setNinex(function() {
        if (Ninex.ninex.commitment_set_time == 0
            || (Ninex.ninex.revealed_time > 0
                && now >= (Ninex.ninex.revealed_time+Ninex.ninex.after_reveal_delay_secs))) {
          // xxx no commitment set
          show('nocommitment');

        } else if (Ninex.ninex.commitment_set_time > 0
                   && Ninex.ninex.revealed_time == 0
                   && (Ninex.ninex.first_guess_time == 0
                       || now < (Ninex.ninex.first_guess_time+Ninex.ninex.guess_window_secs))) {
          Ninex.commitmentEl.innerHTML = Ninex.ninex.commitment;

          // xxx can place a guess
          show('okguess');
          Ninex.change();

        } else if (Ninex.ninex.commitment_set_time > 0
                   && Ninex.ninex.revealed_time == 0
                   && Ninex.ninex.last_guess_time > 0
                   && now >= (Ninex.ninex.last_guess_time+Ninex.ninex.after_last_guess_delay_secs)
                   && now >= (Ninex.ninex.first_guess_time+Ninex.ninex.guess_window_secs)
                   && now < (Ninex.ninex.first_guess_time+Ninex.ninex.reveal_timeout_secs)) {
          // xxx guess window closed, time to reveal
          show('noguess');

        } else if (Ninex.ninex.commitment_set_time > 0
                   && Ninex.ninex.revealed_time == 0
                   && Ninex.ninex.first_guess_time > 0
                   && now >= (Ninex.ninex.first_guess_time+Ninex.ninex.reveal_timeout_secs)) {
          // xxx can timeout
          show('stale');
        }

        window.setTimeout(update, 60000);
      }, function() {
        console.log("error");
      });
    };
    update();
  },
  ninex: null,
  setNinex: function(cb, errcb) {
    var xhr = new XMLHttpRequest();
    xhr.open("GET", Ninex.infoURL, true);
    xhr.onload = function(e) {
      if (xhr.readyState != 4) {
        errcb();
      } else if (xhr.status != 200) {
        errcb();
      } else {
        Ninex.ninex = JSON.parse(xhr.responseText);
        cb();
      }
    };
    xhr.onerror = function (e) {
      errcb();
    };
    xhr.send(null);
  },
  change: function() {
    var digits = Ninex.digitsEl.value;
    if (digits.length == 0) {
      Ninex.stakeinfoEl.style.display = 'none';
    } else {
      Ninex.stakeinfoEl.style.display = 'block';

      var payoutCoeff = 9 * Math.pow(10, digits.length-1) - 1;
      var maxstake = Math.floor(Ninex.ninex.bank / payoutCoeff);
      Ninex.maxstakeEl.innerHTML = maxstake;

      var stake = Ninex.stakeEl.value;
      // xxx check length, check positive integer, check doesn't exceed maxstake
      var payout = stake * (payoutCoeff + 1);
      Ninex.payoutEl.innerHTML = payout;
    }
  },
};
