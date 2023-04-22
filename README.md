# Rotten Haddock

Identify phishing domains from certifiate transparency logs, a rewrite of my previous 'strinkyphish' project utilizing more modern techniques

### I am also trying my first attempt at training a classification model

At this point I've gathered some data (top 1M domains and arond 60,000 previously claffified phishing domains)

This is a refrence for the CSV headers for your training data, as I'm not able to share mine
```
- "td": Task Domain (the specific task or activity that the domain is related to)
- "tld": Top-Level Domain (the highest level of the domain name, such as .com or .org)
- "dl": Domain Length (the total length of the domain name, including subdomains and TLD)
- "nos": Number of Subdomains (the count of subdomains in the domain)
- "nod": Number of Digits (the count of digits in the domain name)
- "noh": Number of Hyphens (the count of hyphens in the domain name)
- "m": Malicious (the target variable indicating whether the domain is a potential phishing attack or not)
```