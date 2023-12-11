<?php
namespace App\Sand;

class Note
{
    public function __construct(string $name)
    {
        $this->name = $name;
        print("<p>executed {$name} Note class constructor</p>");
    }
}