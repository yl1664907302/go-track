<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const messages = ref([])
const size = ref('5')
// 在搜索框输入内容时更新 keyword
const updateKeyword = (event) => {
  size.value = event.target.value
}
// 发送搜索请求
const searchMessages = async () => {
  try {
    const response = await axios.get('/feishu/getmessage', {
      params: {
        index: 'feishu',
        from: 1,
        size: size.value,
        sort_field: 'time.keyword',
        asc: false,
      }
    })
    messages.value = response.data
  } catch (error) {
    console.error('Failed to fetch messages:', error)
    // 如果请求失败，将错误信息置为空数组
    messages.value = []
  }
}
// 在组件挂载时发送初始请求
onMounted(searchMessages)

</script>

<template>
  <!-- 搜索框 -->
  <input type="text" v-model="size" @input="updateKeyword" placeholder="请输入关键字进行搜索">
  <button @click="searchMessages">搜索</button>
  <div>
    <template v-if="messages.length === 0">
      <p>加载中...</p>
    </template>
    <template v-else>
      <ul class="message-list">
        <!-- 遍历 messages 数组，并显示每个消息的内容 -->
        <li v-for="message in messages" :key="message.id" class="message-item">
          <div>
            时间: {{ message.time }}<br>
            群组:    {{message.groupname}}<br>
            平台: {{message.platform}}<br>
          </div>
          <!-- 分割线 -->
          <hr class="divider">
          <!-- 遍历 message.contests 数组，并显示每个项的内容 -->
          <ul>
            <li v-for="(contest, index) in message.contests" :key="index">
              {{ contest.linename }}: {{ contest.linecontext }}
            </li>
          </ul>
        </li>
      </ul>
    </template>
  </div>
</template>

<style scoped lang="less">
/* 添加更多的样式 */

/* 消息列表样式 */
.message-list {
  list-style-type: none;
  padding: 0;
}

/* 消息列表项样式 */
.message-item {
  background-color: rgba(0, 164, 253, 0.34);
  padding: 10px;
  margin-bottom: 5px;
  border-radius: 5px;
  color: #070707; /* 将消息列表项的文本颜色设置为红色 */
}
/* 分割线样式 */
.divider {
  margin-top: 10px;
  border: none;
  border-top: 1px solid #ccc;
}
</style>
