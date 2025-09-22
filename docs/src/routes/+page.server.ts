import fs from 'fs'
import path from 'path'
import type { PageLoad } from './$types'
import type { KeyBinding } from './model.svelte'

const configPath = path.join(process.cwd(), "..", "config.json")
const configJson = fs.readFileSync(configPath, { encoding: 'utf8' })
const configData: KeyBinding = JSON.parse(configJson)

export const load: PageLoad = ({ params }) => {
    return {
        config: configData
    }
}