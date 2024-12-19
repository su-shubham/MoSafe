<script setup lang="ts">
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import TransferCard from './TransferCard.vue'
</script>

<template>
  <div class="min-h-screen bg-background p-8">
    <header class="mb-8">
      <h1 class="heading-primary text-balance">MoSafe Dashboard</h1>
      <p class="text-subtitle text-balance">Your secure financial companion, making money management safe and simple</p>
    </header>

    <div class="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
      <Card>
        <CardHeader>
          <CardTitle class="heading-secondary">Total Balance</CardTitle>
        </CardHeader>
        <CardContent>
          <p class="text-3xl font-semibold font-money">$12,450.80</p>
          <p class="text-subtitle">+2.5% from last month</p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle class="heading-secondary">Recent Transactions</CardTitle>
        </CardHeader>
        <CardContent>
          <p class="text-3xl font-semibold">28</p>
          <p class="text-subtitle">Last 30 days</p>
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle class="heading-secondary">Savings Goal</CardTitle>
        </CardHeader>
        <CardContent>
          <p class="text-3xl font-semibold">75%</p>
          <p class="text-subtitle">$15,000 / $20,000</p>
        </CardContent>
      </Card>
    </div>

    <div class="mt-8">
      <TransferCard />
    </div>
  </div>
</template>
