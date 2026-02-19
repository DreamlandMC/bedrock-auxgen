# bedrock-auxgen

Generate Minecraft Bedrock **AUX IDs** from Mojang’s official item metadata.

`bedrock-auxgen` downloads the Mojang Bedrock item list and converts each item’s `raw_id` into the Bedrock **Aux ID** used internally by the game.

Aux formula:

```
aux = raw_id * 65536
```

---

# ✨ Features

* Uses Mojang’s official Bedrock data source
* Exact AUX calculation
* Outputs ready-to-use JSON mapping
* CLI flags for stdout / custom path
* Deterministic & reproducible
* Perfect for plugins (Nukkit, PocketMine, Geyser)

---

# 📦 Installation

## Go install

```
go install github.com/DreamlandMC/bedrock-auxgen/cmd/bedrock-auxgen@latest
```

Ensure `$GOBIN` is in your PATH.

---

# 🚀 Usage

## Generate default file

```
bedrock-auxgen
```

Creates:

```
typeIdToAux.json
```

---

## Custom output path

```
bedrock-auxgen --out items.json
```

---

## Print mapping to console

```
bedrock-auxgen --stdout
```

Output example:

```
minecraft:acacia_slab: -52822016
minecraft:diamond_sword: 184549376
minecraft:oak_log: 10616832
```

---

## Version

```
bedrock-auxgen --version
```

---

# 📁 Output Format

Generated JSON:

```json
{
  "minecraft:acacia_slab": -52822016,
  "minecraft:diamond_sword": 184549376,
  "minecraft:oak_log": 10616832
}
```

Key = item identifier
Value = AUX ID

---

# 🔌 Java / Plugin Integration

Load once into memory:

```java
private static final Map<String, Long> typeIdToAux;

static {
    try {
        Gson gson = new Gson();
        InputStream stream = MyClass.class.getResourceAsStream("/typeIdToAux.json");
        Type type = new TypeToken<Map<String, Long>>(){}.getType();
        typeIdToAux = gson.fromJson(new InputStreamReader(stream), type);
    } catch (Exception e) {
        throw new RuntimeException(e);
    }
}
```

Lookup:

```java
public static long getAux(String typeId) {
    return typeIdToAux.getOrDefault(typeId, 0L);
}
```

---

# 🧠 Data Source

Mojang Bedrock samples:

https://github.com/Mojang/bedrock-samples

File used:

```
metadata/vanilladata_modules/mojang-items.json
```

---

# ⚙️ How It Works

1. Download Mojang item metadata
2. Parse `data_items`
3. Read `raw_id`
4. Compute AUX:

```
aux = raw_id * 65536
```

5. Export mapping

---

# 🛠 Build From Source

```
git clone https://github.com/DreamlandMC/bedrock-auxgen
cd bedrock-auxgen
go build ./cmd/bedrock-auxgen
```

Run:

```
./bedrock-auxgen
```

---

# 📜 License

MIT

---

# ❤️ Contributions

PRs welcome.
