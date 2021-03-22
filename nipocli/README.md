# CLI
You can use this tool to comunicate and work with nipo

Syntax : `nipocli -t TOKEN -s SERVER -p PORT`

    nipocli -t admin -s 127.0.0.1 -p 2323

## ping
You can be sure that your server is working or not whit ping. If you get the "PONG" everything is fine.

Syntax : `ping`

**Examples**

    nipo > ping
    pong


## status
With this command you can get the slaves status from master

Syntax : `status`

**Examples**

    nipo > status
    { { id : 1 , ip : 127.0.0.1 , status : unhealthy , checkedat : 2021-03-20 19:57:01.614 },{ id : 2 , ip : 127.0.0.1 , status : unhealthy , checkedat : 2021-03-20 19:57:01.614 } }


## set
Which provides you defining your key & value

Syntax : `set KEY VALUE`

**Notes** : 
- The key could be any string without space or tab
- The value could be any string even spaces and tabs, but for reducing the size and increasing the performance, Several spaces or tabs will be concatenated to one space
- The output is set of data with key and value which is the correct amount of stored in memory

**Examples**
	
    nipo > set name My Name       Is  Morteza                    Bashsiz		MB
    {"name":"My Name Is Morteza Bashsiz MB"}
    nipo > set age 30
    {"age":"30"}
    nipo > set sex male
    {"sex":"male"}

## get
which provides you get the value of specific key

Syntax : `get KEY [KEY1 KEY2 KEY3 ... KEYn]`

**Notes** : 
- The key could be any single or multiple string separated with space

**Examples**
	
    nipo > get name
    {"name":"My Name Is Morteza Bashsiz MB"}
    nipo > get name age sex
    {"name":"My Name Is Morteza Bashsiz MB","sex":"male","age":"30"}

## select
which provides you get bulk of specified regex as value

Syntax : `select reg.*`

**Notes** : 
- The key could be any string with standard regex format

**Examples**
	
    nipo > nipo > set my_name Morteza Bashsiz
    {"my_name":"Morteza Bashsiz"}
    nipo > set my_age 30
    {"my_age":"30"}
    nipo > set my_sex male
    {"my_sex":"male"}
    nipo > set your_name Behi Rah
    {"your_name":"Behi Rah"}
    nipo > set your_age 34
    {"your_age":"34"}
    nipo > set your_sex female
    {"your_sex":"female"}
    nipo > get my.*
    nipo > select my.*
    {"my_name":"Morteza Bashsiz","my_age":"30","my_sex":"male"}
    nipo > select your.*
    {"your_name":"Behi Rah","your_age":"34"}
    {"your_sex":"female"}
    nipo > select *.age
    nipo > select .*age.*
    {"your_age":"34","my_age":"30"}
    nipo > 

## sum
which provides you get the sum of values which matches with regex format

Syntax : `sum reg.*`

**Notes** : 
- The key could be any string with standard regex format
- The sum is in float64 format
- If the value of some keys are not numerical it will replace with 0 (zero)

**Examples**
	
    nipo > set f 1.5
    {"f":"1.5"}
    nipo > set fi 2.3
    {"fi":"2.3"}
    nipo > set fir 5 
    {"fir":"5"}
    nipo > set firs 6.7
    {"firs":"6.7"}
    nipo > set first first
    {"first":"first"}
    nipo > sum f.*
    {"f.*":"15.500000"}
    nipo > sum fi.*
    {"fi.*":"14.000000"}
    nipo > sum fir.*
    {"fir.*":"11.700000"}
    nipo > sum firs.*
    {"firs.*":"6.700000"}
    nipo > sum first.*
    {"first.*":"0.000000"}
    nipo >
   
## avg
which provides you get the average of values which matches with regex format

Syntax : `avg reg.*`

**Notes** : 
- The key could be any string with standard regex format
- The sum is in float64 format
- If the value of some keys are not numerical it will replace with 0 (zero)

**Examples**
	
    nipo > set my_age 35.5
    {"my_age":"35.5"}
    nipo > set your_age 30
    {"your_age":"30"}
    nipo > set his_age 23.7
    {"his_age":"23.7"}
    nipo > set her_age 15.2
    {"her_age":"15.2"}
    nipo > avg .*age.*
    {".*age.*":"26.100000"}
    nipo >

## count
which provides you get the count of keys which matches with regex format

Syntax : `count reg.*`

**Notes** : 
- The key could be any string with standard regex format

**Examples**
	
    nipo > count a.*
    {"count":"5432324"}
