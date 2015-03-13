<?php
echo markdown($articleContent);

if ($publishTime){
  echo "<em>First posted: ".date(DATE_RSS, $publishTime)."</em>";
}
