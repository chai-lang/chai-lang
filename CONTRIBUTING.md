## Contributing to ChaiLang

First off, thanks for taking the time to contribute! ❤️
ChaiLang is a community-driven experiment in blending language, culture, and code — and your input makes it better.

### Ways to Contribute

- **Report Bugs** – If something doesn’t work, [open an issue](https://github.com/chai-lang/chai-lang/issues).
- **Suggest Features** – Got an idea for a fun Hindi keyword or syntax sugar? Share it!
- **Improve Docs** – Typos, unclear wording, better examples? All welcome.
- **Contribute Code** – Help improve the lexer, parser, runtime, or anything else.

---

### Getting Started

1. **Fork the repo** and clone your fork:

   ```bash
   git clone https://github.com/your-username/chai-lang.git
   cd chai-lang
   ```

2. **Install Go** (if you haven’t already):
   [https://go.dev/doc/install](https://go.dev/doc/install)

3. **Build the project**:

   ```bash
   go build -o chai-lang ./cmd/chai/main.go
   ```

4. **Run tests**:

   ```bash
   go test ./...
   ```

---

### Code Style

- Use `gofmt` and `go vet`.
- Keep code idiomatic and modular.
- Use meaningful Hindi-English hybrid names where appropriate (`varToken`, `bolStmt`, etc.).

---

### Submitting a Pull Request

1. Create a feature branch:

   ```bash
   git checkout -b my-feature
   ```

2. Make your changes and commit:

   ```bash
   git commit -m "feat: added `agla` keyword support"
   ```

3. Push your branch:

   ```bash
   git push origin my-feature
   ```

4. Open a Pull Request on GitHub. Be descriptive!

---

### Need Help?

If you're stuck or unsure, feel free to open a draft PR or ask in [Discussions](https://github.com/chai-lang/chai-lang/discussions)

---

**Pro Tip:** Make it fun. This language is meant to be playful — like real coding with friends.
