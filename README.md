# digimon-go

A Golang API wrapper for [DAPI](https://digimon-api.com/), a free Digimon API.

## Getting Started

### Install

```bash
go get github.com/stevetoro/digimon-go
```

### Usage

```go
import github.com/stevetoro/digimon-go

func main() {
  c := digimon.NewDigimonClient()
}
```

### Endpoints

digimon-go supports all currently exposed [DAPI](https://digimon-api.com/) endpoints and provides a uniform API for each resource.

<details>
  <summary>Digimon</summary>

  ```go
  // Grab a single Digimon by Name or ID.
  digi, err := c.Digimon.Name("agumon")
  digi, err = c.Digimon.ID(289)

  // Fetch a page of Digimon results and iterate through it.
  page, err := c.Digimon.List()
  page, err = page.Next()
  page, err = page.Prev()

  // Specify query parameters to filter down your Digimon results or jump to a specific page.
  params := resources.DigimonQueryParams{
    Name: "greymon",
    Page: 1,
  }

  page, err = c.Digimon.WithQueryParams(params).List()
  ```
</details>

<details>
  <summary>Attribute</summary>

  ```go
  // Grab a single Attribute by Name or ID.
  attr, err := c.Attribute.Name("vaccine")
  attr, err = c.Attribute.ID(3)

  // Fetch a page of Attribute results and iterate through it.
  page, err := c.Attribute.List()
  page, err = page.Next()
  page, err = page.Prev()

  // Specify query parameters to filter down your Attribute results or jump to a specific page.
  params := resources.QueryParams{
    Name: "data",
    Page: 1,
  }

  page, err = c.Attribute.WithQueryParams(params).List()
  ```
</details>


<details>
  <summary>Field</summary>

  ```go
  // Grab a single Field by Name or ID.
  field, err := c.Field.Name("nightmare soldiers")
  field, err = c.Field.ID(8)

  // Fetch a page of Field results and iterate through it.
  page, err := c.Field.List()
  page, err = page.Next()
  page, err = page.Prev()

  // Specify query parameters to filter down your Field results or jump to a specific page.
  params := resources.QueryParams{
    Name: "night",
    Page: 1,
  }

  page, err = c.Field.WithQueryParams(params).List()
  ```
</details>

<details>
  <summary>Level</summary>

  ```go
  // Grab a single Level by Name or ID.
  level, err := c.Level.Name("armor")
  level, err = c.Level.ID(6)

  // Fetch a page of Field results and iterate through it.
  page, err := c.Level.List()
  page, err = page.Next()
  page, err = page.Prev()

  // Specify query parameters to filter down your Field results or jump to a specific page.
  params := resources.QueryParams{
    Name: "baby",
    Page: 1,
  }

  page, err = c.Level.WithQueryParams(params).List()
  ```
</details>

<details>
  <summary>Type</summary>

  ```go
  // Grab a single Type by Name or ID.
  dType, err := c.Type.Name("cyborg")
  dType, err = c.Type.ID(1)

  // Fetch a page of Type results and iterate through it.
  page, err := c.Type.List()
  page, err = page.Next()
  page, err = page.Prev()

  // Specify query parameters to filter down your Type results or jump to a specific page.
  params := resources.QueryParams{
    Name: "dragon",
    Page: 1,
  }

  page, err = c.Type.WithQueryParams(params).List()
  ```
</details>

<details>
  <summary>Skill</summary>

  ```go
  // Grab a single Skill by Name or ID.
  skill, err := c.Skill.Name("holy jump")
  skill, err = c.Skill.ID(10)

  // Fetch a page of Skill results and iterate through it.
  page, err := c.Skill.List()
  page, err = page.Next()
  page, err = page.Prev()

  // Specify query parameters to filter down your Skill results or jump to a specific page.
  params := resources.QueryParams{
    Name: "holy",
    Page: 1,
  }

  page, err = c.Skill.WithQueryParams(params).List()
  ```
</details>
