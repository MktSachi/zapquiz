<script lang="ts">
   export let quiz: {_id: string, name: string};
</script>

<div role="button" tabindex="0" style="display:flex;align-items:center;justify-content:center;padding:12px 16px;background:#fff;border:2px solid #e6e6e6;border-radius:8px;box-shadow:0 2px 6px rgba(0,0,0,0.06);font-weight:600;color:#222;cursor:pointer;user-select:none;">
    {quiz.name}
</div>
