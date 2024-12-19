<script setup lang="ts">
import { ref } from 'vue'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const amount = ref('')
const recipient = ref('')

const handleTransfer = () => {
  console.log('Transfer:', { amount: amount.value, recipient: recipient.value })
}
</script>

<template>
  <Card class="w-[400px]">
    <CardHeader>
      <CardTitle>Money Transfer</CardTitle>
      <div class="flex flex-col gap-1">
        <span class="text-sm text-muted-foreground">Available Balance</span>
        <span class="text-3xl font-semibold">$2,459.00</span>
      </div>
    </CardHeader>
    <CardContent class="space-y-4">
      <div class="space-y-2">
        <Label for="amount">Amount</Label>
        <Input 
          id="amount"
          v-model="amount"
          placeholder="Enter amount"
          type="number"
        />
      </div>
      
      <div class="space-y-2">
        <Label for="recipient">Recipient</Label>
        <Input 
          id="recipient"
          v-model="recipient"
          placeholder="Enter recipient name"
        />
      </div>

      <Button 
        @click="handleTransfer" 
        class="w-full"
        variant="default"
      >
        Transfer Now
      </Button>
    </CardContent>
  </Card>
</template>
