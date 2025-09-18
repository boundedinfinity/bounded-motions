interface KeyBindingConfig {
    name: string
    key: string
    children: KeyBindingConfig[]
}
import config from './config'

export function init(): KeyBinding {

    const data: KeyBinding = JSON.parse(JSON.stringify(config))

    const walk = (c: KeyBindingConfig, parent?: KeyBinding): KeyBinding => {
        const b = new KeyBinding(c.key, c.name)
        if (parent) parent.add(b)

        if (c.children) {
            for (const child of c.children) {
                walk(child, b)
            }
        }

        return b
    }

    const global = walk(data)

    return global


    // const global = new KeyBinding("", "Bounded Motions")

    // window(global)
    // desktop(global)
    // editor(global)
    // browser(global)

    // return global
}

export class KeyBinding {
    name: string
    key: string
    parent: KeyBinding | undefined
    children: KeyBinding[] = []

    constructor(key: string, name: string) {
        this.key = key
        this.name = name
    }

    duplicate(): KeyBinding {
        return new KeyBinding(this.key, this.name)
    }

    child(key: string, name: string): KeyBinding {
        const child = new KeyBinding(key, name)
        this.add(child)
        return child
    }

    init() {
        for (const child of this.children) {
            child.parent = this
            child.init()
        }
    }

    add(binding: KeyBinding) {
        binding.parent = this
        this.children.push(binding)
    }

    find(...keys: string[]): KeyBinding | undefined {
        let found: KeyBinding | undefined

        if (this.children !== null && this.children !== undefined) {
            for (const key of keys) {
                for (const binding of this.children) {
                    if (binding.key.toLocaleLowerCase() == key.toLowerCase()) {
                        found = binding
                        break
                    }
                }
            }
        }


        return found
    }

    matches(...keys: string[]): boolean {
        return this.find(...keys) != undefined
    }

    path(): KeyBinding[] {
        const found: KeyBinding[] = []
        let current: KeyBinding | undefined = this

        while (current !== undefined) {
            found.push(current)
            current = current.parent
        }

        console.log(found)
        return found.reverse()
    }
}

// ===============================================================================

function os(binding: KeyBinding) {
    const os = binding.child("o", "Operating System")

    const config = os.child("c", "Configuration")
    config.child("a", "Audio")
    const net = config.child("n", "Networking")
    net.child("b", "Bluetooth")
    net.child("w", "Wifi")
    net.child("e", "Hardwired")

    const pkg = os.child("p", "Package")
    pkg.child("a", "Install")
    pkg.child("d", "Remove")
}

function desktop(binding: KeyBinding) {
    const desktop = binding.child("d", "Desktop")

    const focus = desktop.child("f", "Focus")
    directions(focus, { a: true })

    const move = desktop.child("m", "Move")
    directions(move, { a: true })
}

function window(binding: KeyBinding) {
    const window = binding.child("w", "Window")

    const focus = window.child("f", "Focus")
    directions(focus, { a: true })

    const move = window.child("m", "Move")
    directions(move, { a: true })
}

function editor(binding: KeyBinding) {
    const editor = binding.child("e", "Editor")
    editor.child("e", "Explorer")

    const move = editor.child("m", "Move")
    directions(move, { a: true })

    const word = editor.child("w", "Word")
    directions(word, { h: true })
}

function browser(binding: KeyBinding) {
    const browser = binding.child("b", "Browser")

    const tabs = browser.child("t", "Tab")
    directions(tabs, { h: true })

    browser.child("f", "Find in page")
    browser.child("u", "Focus URL bar")
}

function quit(binding: KeyBinding) {
    binding.child("q", "Quit")
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
        binding.child("i", options?.up?.name || "up")
    }

    if (options?.down || options?.a || options?.v) {
        binding.child("k", options?.down?.name || "down")
    }

    if (options?.left || options?.a || options?.h) {
        binding.child("j", options?.left?.name || "left")
    }

    if (options?.right || options?.a || options?.h) {
        binding.child("l", options?.right?.name || "right")
    }
}