

# Proposal: GoYAML — Using Go as its Own Markup Language

## Abstract
In the world of Go development, we often rely on external serialization formats like JSON or YAML to define data and configuration. While these are industry standards, they lack native integration with the Go toolchain, requiring developers to learn separate syntaxes and deal with "bracket hell" or indentation errors. 

This paper proposes **GoYAML**, a methodology and toolset where pure Go is used as the markup language itself. By leveraging the Go compiler and the Language Server Protocol (`gopls`), we can treat data files as first-class Go code. This approach enables native type-safety, seamless refactoring, and complex constraint validation without leaving the Go ecosystem.

---

## 1. Introduction & Background
The project originated during the development of music generation software. While exploring data storage formats, I realized that existing markup languages were unnecessarily complex for certain modeling tasks. Instead of adopting a new Domain Specific Language (DSL), I developed a format based on **pure Go code** to act as a markup language.

### The Core Problem
* **Context Switching:** Moving between Go logic and YAML/JSON data.
* **Lack of Tooling:** Traditional markup doesn't natively support Go-specific refactoring or complex semantic checks.
* **Deep Nesting:** JSON often suffers from deep, unreadable hierarchies.

---

Consider the following mode code:

```go
package models

type Hello struct {
	Name      string
	HelloType HelloType
}
type HelloType string

const (
	Casual HelloType = "Casual"
	Formal HelloType = "Formal"
)

type Country struct {
  Name string
  Hello *Hello
  AlternateHellos []*Hello
}
```

and consider the following data code

```go
```


---

## 2. Technical Architecture
GoYAML splits the data definition into two distinct parts:

* **Model Code:** Defines the structures, types, and constraints.
* **Data Code:** Effectively functions as the "markup" file. 

### Implementation Strategies
1.  **Anonymous Functions:** Data can be defined in anonymous functions. By renaming these to public functions, you can manually initiate data within your code.
2.  **Generated Packages:** The compiler can generate an internal package providing a specialized **API** to initiate data and document the model.

---

## 3. Key Features & Advantages

### A. Zero Learning Curve
If you already program in Go, you don't need to learn a new syntax. The data is Go; the logic is Go.

### B. Native IDE & LSP Support
Because the data files are `.go` files:
* **gopls** naturally checks for syntax and semantic errors.
* **Refactoring:** Renaming a field in your model automatically updates all 1,000+ data files using standard IDE tools.

### C. Versioning & History Mode
GoYAML supports an optional **History Model**. Instead of just storing the final state, the file can contain the creation steps and value assignments (similar to a Git commit history) within the data file itself.

### D. Solving "Bracket Hell"
Unlike JSON, GoYAML encourages a flat structure:
* Every object has a unique identifier (variable name).
* Relationships are established via **pointer assignments** or **slices of pointers**, ensuring the nesting depth rarely exceeds one level.

---

## 4. Modeling Levels (M0 - M3)
For those interested in formal modeling, GoYAML utilizes Go across all four layers of the modeling stack:
* **M3 (Meta-Meta-Model):** The Go language specification itself.
* **M2 (Meta-Model):** A subset of Go used to express the model code.
* **M1 (Model):** The specific domain model defined by the user.
* **M0 (Data):** The actual instances/objects defined in the data files.

---

## 5. Tooling & Performance
The GoYAML compiler generates:
* **Marshaling API:** Provides a stage-based approach to instantiating data.
* **The Probe Program:** A generated utility for basic data navigation and editing tasks.
* **Visual Documentation:** Automatic diagram generation for the model.

**Benchmarks:**
The system has been tested with models containing **100 distinct concepts** and datasets exceeding **50,000 objects** with high stability.

---

## 6. Conclusion
By using Go as a markup language, we bridge the gap between data and logic. We gain the full power of the Go compiler for our data, including type safety, performance, and world-class refactoring tools.

---

**Next Step:** Would you like me to help you expand on the "M0-M3" modeling section to make it more academic, or should I focus on drafting a "Takeaways" section for the GopherCon reviewers?