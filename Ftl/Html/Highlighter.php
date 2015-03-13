<?php
namespace Ftl\Html;

class Highlighter {
  public function php($source){
    $tokens = token_get_all($source);
    $output = '';
    foreach ($tokens as $token){
      if (is_string($token)){
        $output .= htmlEntities($token);
      } else if (is_array($token)) {
        list($id, $text) = $token;
        $name = str_replace('_', '-', strToLower(token_name($id)));
        $text = htmlEntities($text);
        $text = str_replace(
          array(" ",      "\n"),
          array("&nbsp;", "<br/>"),
          $text
        );
        if ($id != T_WHITESPACE){
          $text = "<span class=\"phps-{$name}\">{$text}</span>";
        }
        $output .= $text;
      }
    }
    return $output;
  }
}
