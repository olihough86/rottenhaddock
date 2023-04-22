# Rotten Haddock

Identify phishing domains from certifiate transparency logs, a rewrite of my previous 'strinkyphish' project utilizing more modern techniques

### I am also trying my first attempt at training a classification model

At this point I'm gathering data, I have roghtly 40,000 domains which have been classfied elsewehre, this should be enough for a test data set while I iron out what I'm doing

This is a refrence for the CSV headers
```
    "td": Task Domain (the domain itself)
    "dl": Domain Length (the length of the domain name)
    "nos": Number of Subdomains (the number of subdomains in the domain)
    "nod": Number of Digits (the count of digits in the domain name)
    "noh": Number of Hyphens (the count of hyphens in the domain name)
    "m": Malicious (the target variable indicating whether the domain is phishing or not)
```