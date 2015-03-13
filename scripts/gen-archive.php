<?php
$dir = new DirectoryIterator('posts');

$posts = array();

foreach ($dir as $file){
  if ($file->isDot()) continue;
  $path = $file->getPathname();
  if (!file_exists("{$path}/manifest.json")) continue;
  $manifest = json_decode(file_get_contents("{$path}/manifest.json"));
  $lines = file("{$path}/article.mkd");
  $title = trim(str_replace('_', '\_', trim($lines[0], "#")));
  $postName = $file->getFilename();

  $posts[] = array(
    'path' => '/'.$path,
    'title' => $title,
    'published' => $manifest->published
  );
}

usort($posts, function($a, $b){
  if ($a['published'] == $b['published']) return 0;
  if ($a['published'] > $b['published']) return -1;
  if ($a['published'] < $b['published']) return 1;
});

foreach ($posts as $post){
  $post = (object) $post;

  $date = date('Y-m-d', $post->published);
  echo " * [{$post->title}]({$post->path}) ($date)\n";
}
