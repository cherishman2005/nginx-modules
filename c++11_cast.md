# c++11 cast

## static_cast

enum与int型转换

```
enum class ETileType : uint8
{
ENone, ESurface, ESubsurface, EPlatform, EHole, EEnd
};
```

```
uint8 n = FMath::RandRange(static_cast<int>(ETileType::ESurface), static_cast<int>(ETileType::EEnd) - 1);
ETileType NewTile = static_cast<ETileType>(n);
```
