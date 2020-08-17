Ex: `$ hash-brute-force -w '/home/nix/pass.txt' -h '$2y$10$ojlQW.....'`

Usage:

```
  -c int
    	(Optional) number of cpu core,  -1 = all core (default -1)

  -h string
    	(Required) hash string that need to be found

  -t int
    	(Optional) number of concurrent thread (default 50)
      
  -w string
    	(Required) wordlist file path
```

Example hash string (for 12345) -

`$2y$12$APew2qEmu/1YDnHmdPUT5.idVsU3lN2gE17srB3lC7Jiqsdf2qg9m`

