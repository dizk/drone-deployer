The following is a sample Deployer configuration in your .drone.yml file:

```yaml
deploy:
  deployer:
    task: deploy
    stage: staging

```

If using the common recipe you must set `writable_use_sudo` to false. Ex:

```php
require 'recipe/common.php';

// Drone
set('writable_use_sudo', false);
```
