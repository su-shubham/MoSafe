<script setup lang="ts">
import { ref } from 'vue'
import { Button } from '../components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'

const amount = ref('')
const recipient = ref('')

const handleTransfer = () => {
  console.log('Transfer:', { amount: amount.value, recipient: recipient.value })
}
</script>

<template>
  <Card class="w-full max-w-md">
    <CardHeader>
      <CardTitle class="heading-secondary">MoSafe Transfer</CardTitle>
      <CardDescription class="text-subtitle">Send money securely, anytime, anywhere.</CardDescription>
      <div class="mt-4 p-4 bg-secondary rounded-lg">
        <span class="text-subtitle">Available Balance</span>
        <span class="block text-3xl font-semibold font-money">$2,459.00</span>
      </div>
    </CardHeader>
    <CardContent>
      <div class="space-y-4">
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
          size="lg"
        >
          Transfer Now
        </Button>
      </div>
    </CardContent>
  </Card>
</template>
