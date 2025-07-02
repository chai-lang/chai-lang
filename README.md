# ☕ ChaiLang (Work In Progress)

**ChaiLang** is a fun little programming language made using Go. It’s written in a mix of Hindi + coding, so it feels more desi, more relatable — like coding with your friends over a cup of chai.

---

## What is this?

ChaiLang lets you write programs like this:

```chai
dekh naam hai "muffin";

jab_tak (naam != "thak_gaya") tab_tak {
  bol "Lage raho...";
}
```
### 1. Clone the repository:

```bash 

git clone https://github.com/Dhruva430/chai-lang.git
cd chai-lang
```
### 2. Write a `file_name.chai` file like this:
```bash 
dekh n hai 5;
dekh i hai 1;

jabtak (i <= n) tabtak {
  dekh line hai "";
  dekh space hai 1;
  jabtak (space <= n - i) tabtak {
    line hai line + " ";
    space hai space + 1;
  }

  dekh j hai 1;
  jabtak (j <= i) tabtak {
    line hai line + "* ";
    j hai j + 1;
  }

  bol line;
  i hai i + 1;
}
```
### 3. Run it.
```bash 
./run.sh tokenize file_name.chai
```

Features

Currently supported features:

- `dekh` – declare a variable

- `hai` – assign value

- `bol` – print to output

- `agar` / warna – if/else

- `jab_tak `/ tab_tak – while loop

- `rehne_de` – break out of loop

` haan / nahi` – true/false

- `khali` – null or nil
