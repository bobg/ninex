var Ninex = {
  init: function(infoURL) {
    var update = function() {
      var now = Math.floor(Date.now()/1000); // seconds since unix epoch
      var show = function(name) {
        var el = document.getElementById('initializing');
        el.style.display = 'none';
                 
        el = document.getElementById('nocommitment');
        el.style.display = (name == 'nocommitment') ? 'block' : 'none';

        el = document.getElementById('okguess');
        el.style.display = (name == 'okguess') ? 'block' : 'none';

        el = document.getElementById('noguess');
        el.style.display = (name == 'noguess') ? 'block' : 'none';

        el = document.getElementById('stale');
        el.style.display = (name == 'stale') ? 'block' : 'none';
      };

      var ninex;
      var xhr = new XMLHttpRequest();
      xhr.open("GET", infoURL, true);
      xhr.onload = function(e) {
        if (xhr.readyState === 4) {
          if (xhr.status === 200) {
            var ninex = JSON.parse(xhr.responseText);

            if (ninex.commitment_set_time == 0
                || (ninex.revealed_time > 0
                    && now >= (ninex.revealed_time+ninex.after_reveal_delay_secs))) {
              // xxx no commitment set
              show('nocommitment');

            } else if (ninex.commitment_set_time > 0
                       && ninex.revealed_time == 0
                       && (ninex.first_guess_time == 0
                           || now < (ninex.first_guess_time+ninex.guess_window_secs))) {
              var el = document.getElementById('commitment');
              el.innerHTML = ninex.commitment;

              // xxx can place a guess
              show('okguess');

            } else if (ninex.commitment_set_time > 0
                       && ninex.revealed_time == 0
                       && ninex.last_guess_time > 0
                       && now >= (ninex.last_guess_time+ninex.after_last_guess_delay_secs)
                       && now >= (ninex.first_guess_time+ninex.guess_window_secs)
                       && now < (ninex.first_guess_time+ninex.reveal_timeout_secs)) {
              // xxx guess window closed, time to reveal
              show('noguess');

            } else if (ninex.commitment_set_time > 0
                       && ninex.revealed_time == 0
                       && ninex.first_guess_time > 0
                       && now >= (ninex.first_guess_time+ninex.reveal_timeout_secs)) {
              // xxx can timeout
              show('stale');
            }

            window.setTimeout(update, 60000);
          } else {
            console.error(xhr.statusText);
          }
        }
      };
      xhr.onerror = function (e) {
        console.error(xhr.statusText);
      };
      xhr.send(null);
    };
    update();
  }
};
