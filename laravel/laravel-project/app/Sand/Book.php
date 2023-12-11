<?php
namespace App\Sand;

class Book
{
    public $collectionNumber;
    public $borrowedNumber;
    public $bookedNumber;

    public function __construct()
    {
        print('<p>executed Book class constructor</p>');
    }

    public function setCollectionNumber(int $collectionNumber)
    {
        $this->collectionNumber = $collectionNumber;
    }

    public function setBorrowedNumber(int $borrowedNumber)
    {
        $this->borrowedNumber = $borrowedNumber;
    }

    public function setBookedNumber(int $bookedNumber)
    {
        $this->bookedNumber = $bookedNumber;
    }

    public function getAvailableNumber()
    {
        return $this->collectionNumber - $this->borrowedNumber - $this->bookedNumber;
    }
}