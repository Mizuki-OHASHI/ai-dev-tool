<?php
namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use App\Sand\Note;

class NoteServiceProvider extends ServiceProvider
{
    public function register()
    {
        $this->app->bind('App\Sand\Note', function ($app) {
            $name = "Math3";
            $note = new Note($name);
            return $note;
        });
    }
}