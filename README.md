# bedrock-auxgen

Generate Minecraft Bedrock **AUX IDs** from Mojang’s official item metadata — with optional **custom item support from Resource Packs**.

`bedrock-auxgen` downloads the Mojang Bedrock item list and converts each item’s `raw_id` into the Bedrock **Aux ID** used internally by the game.

Aux formula:

```
aux = item_id << 16
```

For vanilla items:

```
item_id = raw_id 
```

For custom items:

```
item_id = custom_start + index
```

---

# ✨ Features

* Uses Mojang’s official Bedrock data source
* Exact AUX calculation
* Optional custom item AUX generation from Resource Pack
* Configurable custom item start ID
* Deterministic & reproducible
* Outputs ready-to-use JSON mapping
* CLI flags for stdout / custom path
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

## Generate vanilla AUX map

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

Example:

```
minecraft:acacia_slab: -52822016
minecraft:diamond_sword: 184549376
minecraft:oak_log: 10616832
```

---

# 🧩 Custom Items (Resource Pack)

`bedrock-auxgen` can generate AUX IDs for Custom items by reading:

```
<RP>/textures/item_texture.json
```

Example:

```json
{
  "texture_data": {
    "customidentifier:customitem": {
      "textures": "textures/items/customitem"
    }
  }
}
```

---

## Generate with custom items

```
bedrock-auxgen --rp ./MyRP
```

Custom items start at **257** by default.

---

## Change custom start ID

This is useful for third party Serversoftwares because some of them are shifting the IDs a little bit to not run into any issues. (e.g. 10000 For PowerNukkitX):

```
bedrock-auxgen --rp ./MyRP --custom-start 10000
```

Formula:

```
custom_aux = (custom_start + index) << 16
```

---

# 📁 Output Format

Generated JSON:

```json
{
  "minecraft:diamond_sword": 184549376,
  "minecraft:oak_log": 10616832,
  "customidentifier:customitem": 16842752
}
```

Key = item identifier
Value = AUX ID

---

# 🔌 Java / Plugin Integration

Load once:

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

### Vanilla

1. Download Mojang item metadata
2. Read `raw_id`
3. Compute

```
aux = raw_id << 16
```

---

### Custom Items

1. Read `textures/item_texture.json`
2. Collect identifiers
3. Sort alphabetically
4. Assign sequential IDs

```
item_id = custom_start + index
```

5. Compute AUX

```
aux = item_id << 16
```

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
