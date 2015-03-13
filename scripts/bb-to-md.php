#!/usr/bin/env php
<?php

$iterator = new RecursiveIteratorIterator(
  new RecursiveDirectoryIterator('posts')
);

foreach ($iterator as $file){
  $filename = $file->getPathName();
  if (strpos($filename, 'article.mkd') === false) continue;

  file_put_contents($filename, conv($filename));
}

function conv($filename){
  $content = file_get_contents($filename);

  $content = str_replace('[i]', '*', $content);
  $content = str_replace('[/i]', '*', $content);

  $content = str_replace('[b]', '**', $content);
  $content = str_replace('[/b]', '**', $content);

  $content = str_replace('[mono]', '`', $content);
  $content = str_replace('[/mono]', '`', $content);

  $content = preg_replace('/\[link=([^\]]*)\]([^\[]*)\[\/link\]/', '[\2](\1)', $content);

  $content = preg_replace('/\[code\]\w+/', "[code]\n", $content);
  $content = preg_replace('/\w+\[\/code\]/', "\n[/code]", $content);

  $lines = explode("\n", $content);
  $inCode = false;
  foreach ($lines as $k => $line){
    if (trim($line) == '[code]'){
      $lines[$k] = '';
      $inCode = true;
      continue;
    }
    if (trim($line) == '[/code]'){
      $lines[$k] = '';
      $inCode = false;
      continue;
    }
    if ($inCode){
      $lines[$k] = "    {$line}";
    }
  }
  return implode("\n", $lines);
}

