<template>
  <div class="content">
    <h1>Rsbuild with Vue</h1>
    <p>Start building amazing things with Rsbuild.</p>
    <Button label="Submit" />
    
    <div class="card flex justify-center">
        <InputGroup class="sm:!w-96">
            <InputGroupAddon>
                <i class="pi pi-clock"></i>
            </InputGroupAddon>
            <InputGroupAddon>
                <i class="pi pi-star-fill"></i>
            </InputGroupAddon>
            <InputNumber placeholder="Price" />
            <InputGroupAddon>$</InputGroupAddon>
            <InputGroupAddon>.00</InputGroupAddon>
        </InputGroup>
    </div>

  </div>
</template>

<style scoped>
.content {
  display: flex;
  min-height: 100vh;
  line-height: 1.1;
  text-align: center;
  flex-direction: column;
  justify-content: center;
}

.content h1 {
  font-size: 3.6rem;
  font-weight: 700;
}

.content p {
  font-size: 1.2rem;
  font-weight: 400;
  opacity: 0.5;
}
</style>
