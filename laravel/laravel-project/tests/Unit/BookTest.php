<?php

namespace Tests\Unit;

use App\Sand\Book;
use PHPUnit\Framework\TestCase;

class BookTest extends TestCase
{
    /**
     * Test the getAvailableNumber function.
     *
     * @dataProvider availableNumberDataProvider
     */
    public function testGetAvailableNumber($collectionNumber, $borrowedNumber, $bookedNumber, $expectedAvailableNumber): void
    {
        $book = new Book();
        $book->setCollectionNumber($collectionNumber);
        $book->setBorrowedNumber($borrowedNumber);
        $book->setBookedNumber($bookedNumber);

        $actualAvailableNumber = $book->getAvailableNumber();

        $this->assertEquals($expectedAvailableNumber, $actualAvailableNumber);
    }

    /**
     * Data provider for the testGetAvailableNumber function.
     */
    public function availableNumberDataProvider(): array
    {
        return [
            [10, 2, 1, 7],  // Example 1: 10 total, 2 borrowed, 1 booked => 7 available
            [20, 5, 3, 12],  // Example 2: 20 total, 5 borrowed, 3 booked => 12 available
            [15, 3, 2, 10],  // Example 3: 15 total, 3 borrowed, 2 booked => 10 available
            // Add more test cases as needed
        ];
    }
}

