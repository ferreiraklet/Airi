<h1 align="center">Airi</h1> <br>

<p align="center">
  <a href="#--usage--explanation">Usage</a> •
  <a href="#--installation--requirements">Installation</a>
</p>

<h3 align="center">Airi is made for find hidden input parameters in web applications.</h3>


<img src="https://cdn.discordapp.com/attachments/876919540682989609/944639316494274660/unknown.png" align="middle">

## - Installation & Requirements:
```
> git clone https://github.com/ferreiraklet/airi.git

> cd airi

> go build main.go

> mv main airi

> ./airi -h
```
<br>


## - Usage & Explanation:
  Some Web Applications use forms in order to make it stable. Starting from this principle, is possible that the application handle's hidden inputs in source code
  
  Ex: ```<input type="hidden" name="validate" value="test">```
  
  Here it is when Airi appears,
  
  When Web Environment has an input like ```<input type="hidden" name="test" value="">``` and it's value is 0, is very likely the parameter maybe reflected in front end, in this way, making it probably possible to exploit xss reflected.
  
  EXAMPLE:
  
  ```cat index.html```
  output:
  ```<input type="hidden" name="testing" value="">```
  
  Airi reads from stdin
  
  
  <img src="https://cdn.discordapp.com/attachments/876919540682989609/944641808644857876/unknown.png">
  
  You can use a file containing a list of targets as well:
  
  cat targets | airi
  
  
  **Airi only brings to us the url to be tested, so, to test if parameter is reflecting, you can use other tools such as: httpx, kxss, gxss, etc.**


<br>



## This project is for educational and bug bounty porposes only! I do not support any illegal activities!.

If any error in the program, talk to me immediatly.
