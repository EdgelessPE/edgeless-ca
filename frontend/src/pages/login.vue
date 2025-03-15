<template>
  <div
    class="flex flex-col items-center justify-center h-screen bg-gradient-to-b from-#134597 to-#338aca"
  >
    <div
      class="flex items-center justify-between h-lg w-80% bg-#1f1f1f rounded-3xl"
    >
      <div class="flex-1 items-center justify-center flex items-center">
        <div class="h-full text-2xl font-bold">Edgeless CA</div>
      </div>
      <div
        class="w-40% h-full flex flex-col gap-4 rounded-3xl text-#000000ee bg-#eee p-10"
      >
        <div>
          <h2 class="text-2xl font-bold">登录</h2>
          <span>使用 GitHub 注册或登录</span>
        </div>
        <Button class="w-full" label="使用 GitHub 账号" icon="pi pi-github" />
        <Divider align="center" type="dotted">
          <b>OR</b>
        </Divider>
        <span>使用邮箱登录</span>
        <div class="flex flex-col gap-2 w-full">
          <InputText name="email" type="text" placeholder="Email" fluid />
          <Password name="password" placeholder="Password" fluid />
          <div class="text-12px text-#000000ee cursor-pointer" href="#">
            忘记密码？
          </div>
        </div>
        <Button class="w-full" label="Login" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>
