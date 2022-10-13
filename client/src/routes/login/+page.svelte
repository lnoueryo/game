
<div class="layout-container">
  <div class="card">
    <div class="card-body">
      <h1>Login</h1>

      <form
          action="?/login"
          method="POST"
          use:enhance={() => {
            return async ({ result }) => {
              invalidateAll()
              await applyAction(result)
              if(result.type == 'redirect') {
                goto(result.location)
              }
            }
          }}
        >
        <div class="mb-3">
          <label for="username" class="mb-2">Username</label>
          <input id="username" class="form-control form-control-lg" name="username" type="text" placeholder="username" required>
        </div>
        <div class="mb-3">
          <label for="password" class="mb-2">Password</label>
          <input id="password" class="form-control form-control-lg" name="password" type="password" placeholder="password" required>
        </div>
        <div class="message-box">
          {#if form?.invalid}
            <p class="text-danger">Username and password is required.</p>
          {/if}

          {#if form?.credentials}
            <p class="text-danger">You have entered the wrong credentials.</p>
          {/if}
        </div>

        <button type="submit" class="btn btn-primary">Login</button>
      </form>
    </div>
  </div>
</div>

<style>
  .layout-container {
    width: 100%;
    max-width: 450px;
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  .message-box {
    height: 46px;
  }
</style>

<script lang="ts">
  import type { ActionData } from './$types'
  import { applyAction, enhance } from '$app/forms'
  import { goto, invalidateAll } from '$app/navigation'

  export let form: ActionData;

</script>
