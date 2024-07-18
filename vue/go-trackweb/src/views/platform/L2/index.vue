<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const messages = ref([])
const size = ref('1')
const groupname =ref('')
const dateTime=ref('')
const keyword1=ref('')

// 在搜索框输入key时更新 keyword
const mohuGroupname = (event) => {
  groupname.value=event.target.value
}
const mohuSize2 = (event) => {
  size.value = event.target.value
}

const  updateDateTime = (event) => {
  // 在输入更改时更新日期时间变量
  dateTime.value  = event.target.value.replace('T', ' ');
}

const  mohukeyword1 = (event) => {
  // 在输入更改时更新日期时间变量
  keyword1.value  = event.target.value.replace('T', ' ');
}

// 发送模糊搜索请求
const searchMessagesByGroup = async () => {
  try {
    const response = await axios.get('/dingtalk/getmessagemohu', {
      params: {
        index: 'dingtalk',
        from: 1,
        size: size.value,
        sort_field: 'time.keyword',
        asc: false,
        groupname: groupname.value,
        time: dateTime.value,
      }
    })
    messages.value = response.data
    // keyword.value = '' // 重置 keyword 变量为空
  } catch (error) {
    console.error('Failed to fetch messages:', error)
    // 如果请求失败，将错误信息置为空数组
    messages.value = []
  }
}

// 在组件挂载时发送初始请求

onMounted(searchMessagesByGroup)

</script>

<template>

  <div class="search-container">
    <!-- 文本输入框 -->
    组名:<input type="text" v-model.trim="groupname" @input="mohuGroupname" class="number-input" placeholder="">
    数量:<input type="number" v-model.number="size" @input="mohuSize2" class="number-input" placeholder="">
    时间:<input type="datetime-local" v-model="dateTime" @input="updateDateTime" class="datetime-input" placeholder="">
    关键词1:<input type="text" v-model="keyword1" @input="mohukeyword1" class="number-input" placeholder="">
    <!-- 搜索按钮 -->
    <button @click="searchMessagesByGroup" class="search-button">查询</button>
  </div>

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
            群组: {{ message.groupname }}<br>
            平台: {{ message.platform }}<br>
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
/* 消息列表项样式 */
/* 消息列表样式 */
.message-list {
  list-style-type: none; /* 去除默认黑点 */
  padding: 0;
}

/* 消息列表项样式 */
.message-item {
  background-color: #e0e0e0; /* 使用深灰色背景 */
  padding: 20px;
  margin-bottom: 30px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2); /* 加深阴影效果 */
  transform: translateX(20px); /* 向右移动 5px */
}

/* 消息列表项标题样式 */
.message-item h3 {
  margin: 0;
  font-size: 20px;
  font-weight: 600; /* 设置标题加粗 */
  color: #333;
}

/* 消息列表项内容样式 */
.message-item p {
  margin-top: 15px;
  font-size: 16px; /* 提高内容字体大小 */
  color: #666;
}

/* 分割线样式 */
.divider {
  margin-top: 30px;
  border: none;
  border-top: 1px solid #ddd;
}

/*数量框*/
/* 搜索容器样式 */
.search-container {
  display: flex;
  align-items: center;
}

/* 数字输入框样式 */
.number-input {
  padding: 7px 12px;
  border: 2px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
  width: 60px;
  margin-right: 20px; /* 调整输入框与按钮之间的间距 */
  margin-top: 5px;
  margin-left: 20px;
}

/* 搜索按钮样式 */
.search-button {
  padding: 10px 20px;
  background-color: #fd8f00;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.3s ease;
  margin-top: 5px;
}

.search-button:hover {
  background-color: #0056b3;
}

</style>
