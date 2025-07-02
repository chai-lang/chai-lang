# ChaiLang ☕

**ChaiLang** is a work-in-progress, experimental programming language made in Go — designed for the desi coder. It blends Hindi and programming concepts to create a more _relatable_ coding experience — like hacking on something fun with friends over a hot cup of chai.

---

## 🔧 Getting Started

Download the latest release from the [Releases](https://github.com/chai-lang/chai-lang/releases) page.

## 🧪 Try ChaiLang

### 1. A Quick Example

```chai
bol "Hello World!";
```

### 2. Write a File (e.g., `pyramid.chai`)

```chai
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

**Output:**

```
     *
    * *
   * * *
  * * * *
 * * * * *
```

### 3. Run It

**Windows:**

```bash
./chai-lang.exe tokenize pyramid.chai
```

**Linux / macOS:**

```bash
./chai-lang tokenize pyramid.chai
```

---

## 🧠 Language Features

| Hindi Keyword | Purpose                   | Description                      |
| ------------- | ------------------------- | -------------------------------- |
| `dekh`        | Variable Declaration      | Declares a new variable          |
| `hai`         | Assignment                | Assigns a value to a variable    |
| `bol`         | Output                    | Prints output to the console     |
| `agar`        | Conditional (`if`)        | Starts a conditional block       |
| `warna`       | Conditional (`else`)      | Adds an alternate block          |
| `jabtak`      | Loop (`while`)            | Starts a loop                    |
| `tabtak`      | Loop Block                | Marks the loop body              |
| `rehnede`     | Loop Control (`break`)    | Exits the loop                   |
| `agla`        | Loop Control (`continue`) | Skips to next iteration          |
| `haan`        | Boolean (`true`)          | Represents truth                 |
| `nahi`        | Boolean (`false`)         | Represents falsehood             |
| `khali`       | Null                      | Represents a null or empty value |

---

## 🚧 Status

ChaiLang is still under active development. Contributions, ideas, and feedback are always welcome.

---

## ☕ Made with love (and chai)

---
