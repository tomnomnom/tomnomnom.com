<?php
$lines = file('posts.txt');

for ($i = 0; $i < sizeOf($lines); $i = $i + 7){
  $title = trim($lines[$i+1]);
  $body = trim($lines[$i+2]);
  $hardlink = trim($lines[$i+3]);
  $datePublished = trim($lines[$i+4]);
  $dateUpdated = trim($lines[$i+5]);

  $body = str_replace('\r\n', "\n", $body);
  $dateUpdated = str_replace('\N', "null", $dateUpdated);
  $hardlink = str_replace('_', "-", $hardlink);

  mkdir("./tmp/{$hardlink}");
  file_put_contents("./tmp/{$hardlink}/manifest.json", json_encode(array(
    'published' => $datePublished,
    'updated' => $dateUpdated
  )));
  file_put_contents("./tmp/{$hardlink}/article.mkd", "# ".$title."\n\n".$body);
}
