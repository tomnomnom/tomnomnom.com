<?php
$dirs = glob(__DIR__.'/*');
$dirs = array_map(function($dir){
  return str_replace(__DIR__.'/', '', $dir);
}, $dirs);

foreach ($dirs as $dir){
  if (!is_dir($dir)) continue;
  $manifestFile = "{$dir}/manifest.json";
  if (!file_exists($manifestFile)) continue;
  $manifest = json_decode(file_get_contents($manifestFile));

  $article = file("{$dir}/article.mkd");

  $title = str_replace('\\_', '_', trim($article[0], "# \n"));

  file_put_contents("{$dir}/manifest.ini", "title = \"{$title}\"\npublished = {$manifest->published}\nupdated = {$manifest->updated}");
}

