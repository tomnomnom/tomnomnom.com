<!doctype html>
<html lang="en">
  <head>
    <title><?php echo $this->e($title)?: "Tom Hudson's blog - TomNomNom.com"; ?></title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <link href="http://fonts.googleapis.com/css?family=Droid+Serif" rel="stylesheet" type="text/css">
    <link href="http://fonts.googleapis.com/css?family=Droid+Sans+Mono" rel="stylesheet" type="text/css">
    <link href="/styles/main.css" rel="stylesheet" type="text/css"/>
    <link href="/styles/code-dark.css" rel="stylesheet" type="text/css"/>
  </head>
  <body>
    <hgroup>
      <h1><a href="/">TomNomNom.com</a></h1>
      <h2>Tom Hudson's blog</h2>
      <form class="search">
        <input type="text" id="st-search-input" class="st-search-input" />
      </form>
      <div id="st-results-container"></div>
      <script type="text/javascript">
        var Swiftype = window.Swiftype || {};
        (function() {
          Swiftype.key = 'zpxZggzkYvixXZzxd9n7';
          Swiftype.inputElement = '#st-search-input';
          Swiftype.resultContainingElement = '#st-results-container';
          Swiftype.attachElement = '#st-search-input';
          Swiftype.renderStyle = "overlay";

          var script = document.createElement('script');
          script.type = 'text/javascript';
          script.async = true;
          script.src = "//swiftype.com/embed.js";
          var entry = document.getElementsByTagName('script')[0];
          entry.parentNode.insertBefore(script, entry);
        }());
      </script>
    </hgroup>

    <section class="main">
      <?php $this->render($subpage, $subpageData); ?>
    </section>

    <footer>
      <?php if (isSet($lastModified) && $lastModified): ?>
        <div id="lastModified">
          This page last modified: <?php echo date(DATE_RSS, (int) $lastModified); ?>
        </div>
      <?php endif; ?>
    </footer>
  </body>

  <?php if(isSet($gaTracking) && $gaTracking): ?>
    <script type="text/javascript">
      var _gaq = _gaq || [];
      _gaq.push(['_setAccount', 'UA-22278243-3']);
      _gaq.push(['_trackPageview']);

      (function() {
        var ga = document.createElement('script'); ga.type = 'text/javascript'; ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
        var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(ga, s);
      })();
    </script>
  <?php endif; ?>
</html>
