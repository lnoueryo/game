<script lang="ts">
    import { page } from '$app/stores'
    console.log($page)
</script>

<div>
    {$page.data.id}
</div>