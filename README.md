# kulina-test

These are my solutions for the backend developer preliminary test at [Kulina](https://www.kulina.id).

Please see [SOLUTIONS.md](SOLUTIONS.md).

## Basic Concepts

1. Coding
    - Give an example of code smell (1) that usually happened in OOP. Explain and give the preventive action
    - Explain briefly about Dependency Injection, and why is it important?
2. REST API
    - Give 1 example each, of what do’s and don’ts when you handle request while using these method :
        - POST
        - GET

## Basic Coding

Imagine you’re working in a company and you’re told to make a design system and transaction flow. How would you make the most suitable for following users and case:

1. User Side
    - User register
    - User input the address
    - User choose the products to be purchased with subscription and/or one-time purchase scheme
    - User pay the bill
    - User skip the delivery due to certain reasons (ex: They have other agenda that prevent them to receive the delivered goods)
    - User cancel the order
2. Supplier Side
    - Supplier register as seller
    - Supplier create the store and complete the address
    - Supplier create products that can be purchased either daily or one-time purchase
    - Supplier determine the price of each product
    - Supplier determine the selling area
3. Additional:
    - If a product can be sold by more than one seller, define the correct algorithm to determine which order to be sent from which seller (assuming the closest mileage and route)!
    - There is a cut-off time everyday, which is the latest time an order can be placed for the next day delivery. All orders placed beyond cut-off time will automatically delivered on the day after tomorrow.

## Algorithm

1. Nick works at a clothing store. He has a large pile of socks that he must pair by color for sale. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

    For example, there are `n = 7` socks with colors `ar = [1, 2, 1, 2, 1, 3, 2]`. There is one pair of color 1 and one of color 2. There are three odd socks left, one of each color. The number of pairs is 2.

    **Function Description**

    Complete the `sockMerchant` function in the editor below. It must return an integer representing the number of matching pairs of socks that are available.

    `sockMerchant` has the following parameters:

    - `n` : the number of socks in the pile
    - `ar` : the colors of each sock

    **Input Format**

    The first line contains an integer `n` , the number of socks represented in `ar`.  
    The second line contains `n` space-separated integers describing the colors `ar[i]` of the socks in the pile.

    **Constraints**

    - `1 ≤ n ≤ 100`
    - `1 ≤ ar[i] ≤ 100 where 0 ≤ i < n`

    **Output Format**

    Return the total number of matching pairs of socks that Nick can sell.

    **Sample Input**

    ```
    9
    10 20 20 10 10 30 50 10 20
    ```

    **Sample output**

    ```
    3
    ```

2. Bill is an avid hiker. He tracks his hikes meticulously, paying close attention to small details like topography. During his last hike he took exactly `n` steps. For every step he took, he noted if it was an uphill, `U` , or a downhill, `D` step. Gary's hikes start and end at sea level and each step up or down represents a 1 unit change in altitude. We define the following terms:

    - A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
    - A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.

    Given Gary's sequence of up and down steps during his last hike, find and print the number of valleys he walked through.

    For example, if Gary's path is `s = [D D U U U U D D]` , he first enters a valley 2 units deep. Then he climbs out an up onto a mountain 2 units high. Finally, he returns to sea level and ends his hike.

    **Function Description**

    Complete the `countingValleys` function in the editor below. It must return an integer that denotes the number of valleys Gary traversed.

    `countingValleys` has the following parameters:

    - `n` : the number of steps Gary takes
    - `s` : a string describing his path

    **Input Format**

    The first line contains an integer `n` , the number of steps in Gary's hike.  
    The second line contains a single string `s`, of `n` characters that describe his path.

    **Constraints**

    - `2 ≤ n ≤ 106`
    - `s[i] ∈ { U D }`

    **Output Format**

    Print a single integer that denotes the number of valleys Gary walked through during his hike.

    **Sample Input**

    ```
    8
    U D D D U D U U
    ```

    **Sample Output**

    ```
    1
    ```

3. There is an input number: 1.345.679.  
   Write pseudo code (use Golang) that produces following output:

    ```
    1000000
    300000
    40000
    5000
    600
    70
    9
    ```

4. Andrew walks through 100 switches from point A to point B with 1 to 100 as the number. At the first trip, Andrew presses all of the switches so lamps are on. Second trip, Andrew only presses switches that multiplying of two. The next trip, Andrew presses switches that multiplying of three. Andrew repeats his trips from A to B for 100 times. Write down the code to determine how many lamps will turn on after 100 trips from A to B.
