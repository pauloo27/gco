# GCO

Write God-tier commits.

![CI status](https://ci.notagovernment.agency/api/v1alpha/badges/442b5a0e-cfb1-4ce0-affa-6eb59053bd4f)

<img src="https://github.com/Pauloo27/gommit/raw/master/.assets/commit.gif" alt="Usage example gif" height="250x" />

## Hooks

You can add Pre and Post commit hooks in `.gcorc.json`, like that:
```json
...

  "PreCommit": [
    {
      "Command": "make",
      "ExitOnFail": true,
      "Ask": false
    }
  ],
  "PostCommit": [
    {
      "Command": "git push",
      "ExitOnFail": true,
      "Ask": true
    }
  ]
  
...
```

A hook is made of:
- Command: the command to run.
- ExitOnFail: if the command fail, should the commit be cancelled?
- Ask: ask for confirmation before running the hook.

_GCO hooks does not replace Git hooks._

## License

<img src="https://i.imgur.com/AuQQfiB.png" alt="GPL Logo" height="100px" />

This project is licensed under [GNU General Public License v2.0](./LICENSE).

This program is free software; you can redistribute it and/or modify 
it under the terms of the GNU General Public License as published by 
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

