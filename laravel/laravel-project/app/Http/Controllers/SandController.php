<?php

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\DB;

use App\Sand\Book;
use App\Sand\Magazine;
use App\Sand\Note;

class SandController extends Controller
{
    public function newBook()
    {
        $book = new Book();
        return '<p>executed newBook method</p>';
    }

    public function newBook2()
    {
        $book = resolve('App\Sand\Book');
        return '<p>executed newBook2 method</p>';
    }

    public function newBook3(Book $book)
    {
        return '<p>executed newBook3 method</p>';
    }

    private $magazine;
    public function __construct(Magazine $magazine)
    {
        $this->magazine = $magazine;
    }

    public function newMathNote(Note $note3)
    {
        $note1 = new Note('Math1');
        $note2 = resolve('App\Sand\Note');
        return '<p>executed newNote method</p>';
    }

    public function runningRawSQLQueries()
    {
        $sql = 'SELECT * FROM user WHERE created_at > :date';

        $bindParams = [':date' => '2023-11-01 00:00:00'];
        $users = DB::select($sql, $bindParams);

        $table = '<table>';
        $table .= '<tr><th>Name</th><th>Email</th></tr>';
        foreach ($users as $user) {
            $table .= '<tr>';
            $table .= '<td>' . $user->name . '</td>';
            $table .= '<td>' . $user->email . '</td>';
            $table .= '</tr>';
        }
        $table .= '</table>';

        return $table;
    }
}