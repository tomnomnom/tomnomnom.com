<?php
namespace Ftl\Html;

class Renderer {
  public function render($file, array $data = array()){
    extract($data);
    require $file;
  }

  public function e($data){
    return htmlspecialchars($data, ENT_QUOTES, 'UTF-8');
  }
}
