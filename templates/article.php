<?php
echo markdown($articleContent);

if ($publishTime){
  echo "<em class=\"postTime\">First posted: ".date(DATE_RSS, $publishTime)."</em>";
}
