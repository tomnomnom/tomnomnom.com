<?php
header("Content-type: text/html; charset=utf-8");
date_default_timezone_set('UTC');

// Hacky but functional
define('IN_PRODUCTION', (bool) ($_SERVER['SERVER_NAME'] === 'tomnomnom.com'));
if (!IN_PRODUCTION){
  error_reporting(-1);
  ini_set('display_errors', 'on');
  ini_set('html_errors', 'on');
}
define('ROOT_DIR', realpath(__DIR__.'/../'));

// Markdown won't autoload at the moment; no namespaces etc
require ROOT_DIR.'/Ftl/Html/Markdown.php';

spl_autoload_register(function($class){
  $class = str_replace('\\', '/', ltrim($class, '\\'));
  require ROOT_DIR."/{$class}.php";
});

set_exception_handler(function($e){
  switch ($e->getCode()){
    case 404:
      header("HTTP/1.1 404 Not Found");
      break;
  }
  $r = new \Ftl\Html\Renderer();
  $r->render(ROOT_DIR.'/templates/main.php', array(
    'title'        => 'Oh noes! Waht you doen?!',
    'subpage'      => ROOT_DIR.'/templates/error.php',
    'subpageData'  => array(
      'exception' => $e
    ) 
  ));
});

/********/
/* Meat */
/********/

$renderer = new \Ftl\Html\Renderer();

$request = parse_url($_SERVER['REQUEST_URI']);
$router = new \Ftl\Http\Router($request['path']);
$router->addRoute(
  array(
    '/', 
    '/posts', 
    '/posts/:article'
  ), 
  function($article = 'index') use($renderer){
    $articleFile = ROOT_DIR."/posts/{$article}/article.mkd";
    if (!file_exists($articleFile)){

      // Try old-style links 
      if (strpos($article, '_') !== false){
        $article = str_replace('_', '-', $article);
        header("HTTP/1.1 301 Moved Permanently"); 
        header("Location: /posts/{$article}"); 
        return false;
      }

      throw new \Exception("That is not an article I have written.", 404);
    }
    $articleContent = file_get_contents($articleFile);
    $lastModified = filemtime($articleFile);

    $manifestFile = ROOT_DIR."/posts/{$article}/manifest.ini";
    $title = null;
    if (file_exists($manifestFile)){
      $manifest = parse_ini_file($manifestFile);
      $title = isSet($manifest['title'])? $manifest['title'].' - TomNomNom.com' : null;
    }

    $renderer->render(ROOT_DIR.'/templates/main.php', array(
      'title'        => $title,
      'subpage'      => ROOT_DIR.'/templates/article.php',
      'lastModified' => $lastModified,
      'gaTracking'   => IN_PRODUCTION, //Only track in production
      'subpageData'  => array(
        'articleContent' => $articleContent,
        'publishTime' => isSet($manifest['published'])? $manifest['published'] : null
      ) 
    ));
  }
);

$router->dispatch();


