| Category |     |     |     |     | Description                        |
| -------- | --- | --- | --- | --- | ---------------------------------- |
| Root     | w   | m   | i   |     | Window -> move -> up               |
| Window   |     |     | j   |     | Window -> move -> left             |
| Window   |     |     | k   |     | Window -> move -> down             |
| Window   |     |     | l   |     | Window -> move -> right            |
| Window   |     | f   | i   |     | Window -> focus -> up              |
| Window   |     |     | j   |     | Window -> focus -> left            |
| Window   |     |     | k   |     | Window -> focus -> down            |
| Window   |     |     | l   |     | Window -> focus -> right           |
| Window   |     | x   |     |     | Window -> close                    |
| Window   |     | r   | a   | i   | Window -> resize -> all -> larger  |
| Window   |     |     |     | k   | Window -> resize -> all -> smaller |
| Window   |     |     | i   | i   | Window -> resize -> top -> up      |
| Window   |     |     |     | k   | Window -> resize -> top -> down    |
| Window   |     |     | k   | i   | Window -> resize -> bottom -> up   |
| Window   |     |     |     | k   | Window -> resize -> bottom -> down |
| Window   |     |     | j   | j   | Window -> resize -> left -> left   |
| Window   |     |     |     | l   | Window -> resize -> left -> right  |
| Window   |     |     | l   | j   | Window -> resize -> right -> left  |
| Window   |     |     |     | l   | Window -> resize -> right -> right |

```
├── w
    ├── m
    │   ├── i       move window up
    │   ├── j       move window left
    │   ├── k       move window down
    │   └── l       move window right
    └── f
    │   ├── i       focus window up
    │   ├── j       focus window left
    │   ├── k       focus window down
    │   └── l       focus window left
    ├── x           close window
    └── r
        ├── a
        │   ├── i   resize window larger
        │   └── k   resize window smaller
        ├── i
        │   ├── i   move top of window up
        │   └── k   move top of window down
        ├── k
        │   ├── i   move bottom of window up
        │   └── k   move bottom of window down
        ├── j
        │   ├── j   move left of window left
        │   └── l   move left of window right
        └── l
            ├── j   move right of window left
            └── l   move right of window right

```
