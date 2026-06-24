---
title: "TheGoldEconomy"
logo: "/tge-logo.png"
subtitle: " A simple gold-based economy plugin"
accent_color: "#fac900"
toc: true

badges:
  - alt: "Modrinth Downloads"
    src: "https://img.shields.io/modrinth/dt/thegoldeconomy?style=for-the-badge&labelColor=%23555555&color=%23fac900"
  - alt: "Spiget Rating"
    src: "https://img.shields.io/spiget/rating/102242?style=for-the-badge&labelColor=%23555555&color=%23fac900"
  - alt: "Spiget tested server versions"
    src: "https://img.shields.io/spiget/tested-versions/102242?style=for-the-badge&labelColor=%23555555&color=%23fac900"
    
buttons:
  - text: "Download on Modrinth"
    link: "https://modrinth.com/plugin/thegoldeconomy"
    button: true
  - text: "View on GitHub"
    link: "https://github.com/confusedalex/GoldEconomy"
---
## Introduction

TheGoldEconomy is a powerful economy plugin that allows servers to
manage their gold-based currency through a bank system. Players can
deposit and withdraw gold and send money to each other. With support
for both gold nuggets, ingots and raw gold this plugin is designed to
be lightweight, easy to configure, and fully compatible with popular
plugins like Vault and Towny.

## Features

- Different gold translation modes
  - Nugget Mode: 1 Nugget = 1$
  - Ingot Mode: 1 Ingot = 1$
  - Raw Gold Mode: 1 Raw Gold = 1$
- Extensive plugin comatability
  - Vault support to work with other plugins
  - PlaceholderAPI
  - Towny bank plots
- Option to completly remove gold drops from mobs
- Many languages included
- No database required

## Dependencies

-  **[Vault](https://www.spigotmc.org/resources/vault.34315/)**

## Additionals Information

<details>
  <summary>Commands</summary>

- **/bank balance | /balance | /bal**  
  Displays your current bank balance (e.g, `/bank balance`).
- **/bank balance <player>**  
  Shows the balance of the specified player (e.g., `/bank balance Steve`).
- **/bank deposit <gold>**  
  Deposits the specified amount of gold from your inventory into your bank account (e.g., `/bank deposit 10`). To deposit everything use `/bank deposit` without an amount.
- **/bank withdraw <gold>**  
  Withdraws the specified amount of gold from your bank account into your inventory (e.g., `/bank withdraw 5`). To withdraw everything use `/bank withdraw` without an amount.
- **/bank pay <player> <gold>**  
  Transfers the specified amount of gold to another player (e.g., `/bank pay Alex 20`).
  
</details>

<details>
  <summary>Permissions</summary>

| Permission                    | Command                | default |
|-------------------------------|------------------------|---------|
| thegoldeconomy.balance        | /bank balance          | yes     |
| thegoldeconomy.balance.others | /bank balance <player> | yes     |
| thegoldeconomy.deposit        | /bank deposit          | yes     |
| thegoldeconomy.withdraw       | /bank withdraw         | yes     |
| thegoldeconomy.pay            | /bank pay              | yes     |
| thegoldeconomy.set            | /bank set              | no      |
| thegoldeconomy.add            | /bank add              | no      |
| thegoldeconomy.remove         | /bank remove           | no      |

  
</details>

<details>
<summary>Placeholders</summary>

The following placeholders are available if using PlaceholderAPI

- `thegoldeconomy_inventoryBalance`
- `thegoldeconomy_bankbalance`
- `thegoldeconomy_totalBalance`

</details>

<details>
<summary>Config File</summary>

```yaml
# Remove Gold Drops from Mobs like Piglins? (default: true)
removeGoldDrop: true
# Should the plugin check for updates? (default: true)
updateCheck: true
# Valid language are:
# Brazilian Portuguese: pt_BR
# Bulgarian: bg_BG
# English: en_US
# German: de_DE
# Japanese: jp_JP
# Norwegian: nb_NO
# Polish: pl_PL
# Simplified Chinese: zh_CN
# Spanish: es_ES
# Tamil: ta
# Turkish: tr_TR
# Ukrainian: uk
# Russian: ru
language: "en_US"
# Do you want to restrict bank commands to bank plots (requires Towny)
restrictToBankPlot: false
# Prefix
prefix: "TheGoldEconomy"
# This value sets the base domination of the economy
# 'nuggets' = 1 nugget is 1 currency, 1 ingot is 9, 1 block is 81
# 'ingots'  = 1 ingots is 1 currency, 1 block is 9
# 'raw' = 1 raw gold is 1 currency, 1 block is 9
base: "nuggets"
```
</details>
