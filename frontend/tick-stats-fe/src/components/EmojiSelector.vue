<template>
    <div class="emoji-container">
        
        <div
          v-for="emoji in emojis"
          :key="emoji"
          class="emoji-item"
        >
          <v-btn @click="selectEmoji(emoji)" style="font-size: 24px;" variant="plain">
            {{ emoji }}
          </v-btn>
        </div>
      </div>
</template>

<script setup>
import { ref } from 'vue';

const emojis = ref([
  "😋", "🥰", "🤓", "🥳", "😊", "😕", "⭐️", "🤩", "🙂", "😀", "😌", "😎", "🥳", "🤯", "🌟", "💫", "🌎", "👋", "🚀", "🛸", "😄", "😂", "❤️"
]);


</script>

<script>
export default {
    methods: {
        selectEmoji(emoji) {
            this.$emit('select', emoji)
        }
    }
}
</script>

<style>

.emoji-container {
    display: flex;
    overflow-x: auto;
    white-space: nowrap;
  }
  
  .emoji-item {
    flex: 0 0 auto;
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 8px;
  }
</style>