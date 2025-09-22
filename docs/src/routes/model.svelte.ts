interface KeyBindingConfig {
    name: string
    key: string
    expanded?: boolean
    children: KeyBindingConfig[]
}

export function init(data:KeyBinding): KeyBinding {
    const walk = (c: KeyBindingConfig, parent?: KeyBinding): KeyBinding => {
        const b = new KeyBinding(c.key, c.name, { expanded: c.expanded })
        if (parent) parent.addChild(b)

        if (c.children) {
            for (const child of c.children) {
                walk(child, b)
            }
        } else {
            c.children = []
        }

        if(c.expanded) {
            b.walkUp(function(cur:KeyBinding): boolean {
                cur.expanded = true
                return true
            })
        }

        return b
    }

    const global = walk(data)

    return global
}

export class KeyBinding {
    name: string
    key: string
    parent: KeyBinding | undefined
    expanded = $state(false)
    children: KeyBinding[] = []

    constructor(key: string, name: string, options?: { expanded?: boolean }) {
        this.key = key
        this.name = name
        this.expanded = options?.expanded || false
    }

    toggleExpended() {
        this.expanded = !this.expanded
    }

    duplicate(): KeyBinding {
        return new KeyBinding(this.key, this.name)
    }

    newChild(key: string, name: string): KeyBinding {
        const child = new KeyBinding(key, name)
        this.addChild(child)
        return child
    }

    init() {
        for (const child of this.children) {
            child.parent = this
            child.init()
        }
    }

    addChild(binding: KeyBinding) {
        binding.parent = this
        this.children.push(binding)
    }

    walkDown(fn: (binding: KeyBinding) => boolean): boolean {
        if (fn === undefined || fn === null)
            return false


        if (!fn(this)) {
            return false
        }

        if (this.children === null && this.children === undefined) {
            return true
        }

        for (const child of this.children) {
            if (!child.walkDown(fn)) {
                return false
            }
        }

        return true
    }

    walkUp(fn: (binding: KeyBinding) => boolean): boolean {
        if (fn == undefined || fn == null)
            return false

        if (!fn(this)) {
            return false
        }

        if (this.parent === undefined || this.parent === null) {
            return false
        }

        return this.parent.walkUp(fn)
    }

    find(...keys: string[]): KeyBinding | undefined {
        let found: KeyBinding | undefined

        function fn(key: string): (binding: KeyBinding) => boolean {
            return function (binding: KeyBinding): boolean {
                if (binding.key.toLocaleLowerCase() == key.toLowerCase()) {
                    found = binding
                    return false
                }
                return true
            }
        }

        for (const key of keys) {
            this.walkDown(fn(key))
        }

        return found
    }

    matches(...keys: string[]): boolean {
        return this.find(...keys) != undefined
    }

    path(): KeyBinding[] {
        const found: KeyBinding[] = []
        const fn = (binding: KeyBinding): boolean => {
            found.push(binding)
            return true
        }

        this.walkUp(fn)

        return found.reverse()
    }
}

// ===============================================================================

function os(binding: KeyBinding) {
    const os = binding.newChild("o", "Operating System")

    const config = os.newChild("c", "Configuration")
    config.newChild("a", "Audio")
    const net = config.newChild("n", "Networking")
    net.newChild("b", "Bluetooth")
    net.newChild("w", "Wifi")
    net.newChild("e", "Hardwired")

    const pkg = os.newChild("p", "Package")
    pkg.newChild("a", "Install")
    pkg.newChild("d", "Remove")
}

function desktop(binding: KeyBinding) {
    const desktop = binding.newChild("d", "Desktop")

    const focus = desktop.newChild("f", "Focus")
    directions(focus, { a: true })

    const move = desktop.newChild("m", "Move")
    directions(move, { a: true })
}

function window(binding: KeyBinding) {
    const window = binding.newChild("w", "Window")

    const focus = window.newChild("f", "Focus")
    directions(focus, { a: true })

    const move = window.newChild("m", "Move")
    directions(move, { a: true })
}

function editor(binding: KeyBinding) {
    const editor = binding.newChild("e", "Editor")
    editor.newChild("e", "Explorer")

    const move = editor.newChild("m", "Move")
    directions(move, { a: true })

    const word = editor.newChild("w", "Word")
    directions(word, { h: true })
}

function browser(binding: KeyBinding) {
    const browser = binding.newChild("b", "Browser")

    const tabs = browser.newChild("t", "Tab")
    directions(tabs, { h: true })

    browser.newChild("f", "Find in page")
    browser.newChild("u", "Focus URL bar")
}

function quit(binding: KeyBinding) {
    binding.newChild("q", "Quit")
}

interface kbd {
    name?: string
    description?: string
}

function directions(binding: KeyBinding, options?: {
    a?: boolean, h?: boolean, v?: boolean,
    up?: kbd, down?: kbd, left?: kbd, right?: kbd
}) {
    if (options?.up || options?.a || options?.v) {
        binding.newChild("i", options?.up?.name || "up")
    }

    if (options?.down || options?.a || options?.v) {
        binding.newChild("k", options?.down?.name || "down")
    }

    if (options?.left || options?.a || options?.h) {
        binding.newChild("j", options?.left?.name || "left")
    }

    if (options?.right || options?.a || options?.h) {
        binding.newChild("l", options?.right?.name || "right")
    }
}