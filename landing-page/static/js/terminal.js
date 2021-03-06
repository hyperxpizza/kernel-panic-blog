var util = util || {};
util.toArray = function(list) {
  return Array.prototype.slice.call(list || [], 0);
};

var Terminal = Terminal || function(cmdLineContainer, outputContainer) {
  window.URL = window.URL || window.webkitURL;
  window.requestFileSystem = window.requestFileSystem || window.webkitRequestFileSystem;

  var cmdLine_ = document.querySelector(cmdLineContainer);
  var output_ = document.querySelector(outputContainer);

  const CMDS_ = [
    'cat', 'clear', 'date', 'help', 'uname', 'whoami', 'ls', 'ifconfig', 'pwd', 
  ];

  const DIRS_ = [
    'README.md', 'blog.sh'
  ];
  
  var fs_ = null;
  var cwd_ = null;
  var history_ = [];
  var histpos_ = 0;
  var histtemp_ = 0;

  window.addEventListener('click', function(e) {
    cmdLine_.focus();
  }, false);

  cmdLine_.addEventListener('click', inputTextClick_, false);
  cmdLine_.addEventListener('keydown', historyHandler_, false);
  cmdLine_.addEventListener('keydown', processNewCommand_, false);

  //
  function inputTextClick_(e) {
    this.value = this.value;
  }

  //
  function historyHandler_(e) {
    if (history_.length) {
      if (e.keyCode == 38 || e.keyCode == 40) {
        if (history_[histpos_]) {
          history_[histpos_] = this.value;
        } else {
          histtemp_ = this.value;
        }
      }

      if (e.keyCode == 38) { // up
        histpos_--;
        if (histpos_ < 0) {
          histpos_ = 0;
        }
      } else if (e.keyCode == 40) { // down
        histpos_++;
        if (histpos_ > history_.length) {
          histpos_ = history_.length;
        }
      }

      if (e.keyCode == 38 || e.keyCode == 40) {
        this.value = history_[histpos_] ? history_[histpos_] : histtemp_;
        this.value = this.value; // Sets cursor to end of input.
      }
    }
  }

  //
  function processNewCommand_(e) {

    if (e.keyCode == 9) { // tab
      e.preventDefault();
      // Implement tab suggest.
    } else if (e.keyCode == 13) { // enter
      // Save shell history.
      if (this.value) {
        history_[history_.length] = this.value;
        histpos_ = history_.length;
      }

      // Duplicate current input and append to output section.
      var line = this.parentNode.parentNode.cloneNode(true);
      line.removeAttribute('id')
      line.classList.add('line');
      var input = line.querySelector('input.cmdline');
      input.autofocus = false;
      input.readOnly = true;
      output_.appendChild(line);

      if (this.value && this.value.trim()) {
        var args = this.value.split(' ').filter(function(val, i) {
          return val;
        });
        var cmd = args[0].toLowerCase();
        args = args.splice(1); // Remove cmd from arg list.
      }

      console.log(input.value)

      switch (input.value) {
        case 'cat README.md':
          output('<div class="readme"># kernel-panic.pl<br>Coming soon<br><a href="https://github.com/hyperxpizza/kernel-panic-blog" style="text-decoration: none; color: green;"> https://github.com/hyperxpizza/kernel-panic-blog</a><br>hyperxpizza2gmail.com</div>');
          break;
        case 'cat':
          output('<div><b>NAME</b><br><p style="padding-left: 5rem;">cat - concatenate files and print on the standard output</p><br><b>SYNOPSIS</b><br><p style="padding-left: 5rem;">cat [OPTION]... [FILE]...</p></div>');
          return
        case 'clear':
          output_.innerHTML = '';
          this.value = '';
          return;
        case 'date':
          output( new Date() );
          break;
        case 'help':
          output('<div class="ls-files">' + CMDS_.join('<br>') + '</div>');
          break;
        case 'uname':
          output(navigator.appVersion);
          break;
        case 'whoami':
          $.getJSON('https://ipinfo.io', function(data){
            output('<p>ip: ' + data.ip + '<br>country: ' + data.country + '<br>region: ' + data.region + '<br>city: ' + data.city + '<br>location: ' + data.loc + '<br>timezone: ' + data.timezone + '</p>');
          });
          break;
        case 'ls':
          output('<div class="ls-files">' + DIRS_.join('<br>') + '</div>');
          break;
        case 'ls -la':
          output('<p>drwxrwxr-x  2 user user 4096 sep 16 16:37 . </p> <p>drwxrwxr-x  17 user user 4096 sep 16 16:37 .. </p> <p>drwxrwxr-x  1 user user 116 sep 16 16:37 README.md </p> <p>drwxrwxr-x  1 user user 39 sep 16 16:37 blog.sh </p>');
          break;
        case 'pwd':
          output('/home/user');
          break;
        case 'ifconfig':
          $.getJSON('https://ipinfo.io', function(data){
            output('<div> eth1: flags=4163 UP,BROADCAST,RUNNING,MULTICAST  mtu 1500<br><div style="padding-left:5em;">inet' + data.ip + 'netmask 255.255.240.0  broadcast 46.41.143.255<br>inet6 fe80::8dd:cff:fe1a:1486  prefixlen 64  scopeid 0x20<link><br>ether 00:1c:42:64:ed:a2  txqueuelen 1000  (Ethernet)<br>RX packets 43741455  bytes 5474143689 (5.4 GB)<br>RX errors 0  dropped 427522  overruns 0  frame 0 <br>TX packets 368890  bytes 56717418 (56.7 MB)<br>TX errors 0  dropped 4 overruns 0  carrier 0  collisions 0</div></div><div>lo: flags=73 UP,LOOPBACK,RUNNING  mtu 65536<br><div style="padding-left: 5em";>inet 127.0.0.1  netmask 255.0.0.0 <br> inet6 ::1  prefixlen 128  scopeid 0x10 host <br>loop  txqueuelen 1000  (Local Loopback) <br>RX packets 578  bytes 215105 (215.1 KB) <br>RX errors 0  dropped 0  overruns 0  frame 0 <br>TX packets 578  bytes 215105 (215.1 KB) <br>TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0 <br></div></div>');
          });
          break;
        case './blog.sh':
          output('<p>[+] Coming soon: kernel-panic.pl</p>');
          break;
        case 'blog.sh':
          output('<p>bash: ./blog.sh: Permission denied</p>');
          break;
        default:
          if (input.value) {
            output(input.value + ': command not found');
          }
      };

      window.scrollTo(0, getDocHeight_());
      this.value = ''; // Clear/setup line for next input.
    }
  }

  //
  function formatColumns_(entries) {
    var maxName = entries[0].name;
    util.toArray(entries).forEach(function(entry, i) {
      if (entry.name.length > maxName.length) {
        maxName = entry.name;
      }
    });

    var height = entries.length <= 3 ?
        'height: ' + (entries.length * 15) + 'px;' : '';

    // 12px monospace font yields ~7px screen width.
    var colWidth = maxName.length * 7;

    return ['<div class="ls-files" style="-webkit-column-width:',
            colWidth, 'px;', height, '">'];
  }

  //
  function output(html) {
    output_.insertAdjacentHTML('beforeEnd', '<p>' + html + '</p>');
  }

  // Cross-browser impl to get document's height.
  function getDocHeight_() {
    var d = document;
    return Math.max(
        Math.max(d.body.scrollHeight, d.documentElement.scrollHeight),
        Math.max(d.body.offsetHeight, d.documentElement.offsetHeight),
        Math.max(d.body.clientHeight, d.documentElement.clientHeight)
    );
  }

  //
  return {
    init: function() {
      output('<h2 style="letter-spacing: 4px">kernel-panic.pl web terminal</h2><p>' + new Date() + '</p><p>Enter "help" for more information.</p>');
    },
    output: output
  }
};

