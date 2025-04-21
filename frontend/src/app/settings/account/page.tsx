"use client"

import { useState, useEffect } from 'react'
import IndexLayout from '@/app/layouts/index-layout'
import { ThemeToggle } from '@/components/theme-toggle'

interface User {
  id: string
  name: string
  email: string
  role: string
  provider: string
  createdAt: string
  updatedAt: string
}

export default function UserSettings() {
  const [user, setUser] = useState<User | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const accessToken = localStorage.getItem('auth_token')
        if (!accessToken) {
          throw new Error('Access token is missing')
        }

        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/user/me`, {
          headers: {
            'Authorization': `Bearer ${accessToken}`,
          },
        })

        if (!res.ok) {
          throw new Error('Failed to fetch user data')
        }

        const responseData = await res.json()
        console.log(responseData)

        if (responseData.data) {
          setUser(responseData.data.user)
        } else {
          throw new Error('User data not found in the response')
        }
      } catch (error) {
        setError("An error occurred. Please try again." + error)
      } finally {
        setLoading(false)
      }
    }

    fetchUserData()
  }, [])

  if (loading) {
    return <p>Loading...</p>
  }

  if (error) {
    return <p>Error: {error}</p>
  }

  if (!user) {
    return <p>No user data available.</p>
  }

  return (
    <IndexLayout>
      <div className="flex min-h-screen w-full items-center justify-center p-6 md:p-10">
        <div className="w-full max-w-lg space-y-4">
          <h1 className="text-3xl font-semibold">User Settings</h1>
          <div className="bg-white p-6 rounded-md shadow-md space-y-4">
            <div>
              <p className="text-lg font-semibold">Name:</p>
              <p>{user.name}</p>
            </div>
            <div>
              <p className="text-lg font-semibold">Email:</p>
              <p>{user.email}</p>
            </div>
            <div>
              <p className="text-lg font-semibold">Role:</p>
              <p>{user.role}</p>
            </div>
            <div>
              <p className="text-lg font-semibold">Provider:</p>
              <p>{user.provider}</p>
            </div>
            <div>
              <p className="text-lg font-semibold">Account Created At:</p>
              <p>{user.createdAt ? new Date(user.createdAt).toLocaleDateString() : 'N/A'}</p>
            </div>

            <div>
              <p className="text-lg font-semibold">Last Updated At:</p>
              <p>{user.updatedAt ? new Date(user.updatedAt).toLocaleDateString() : 'N/A'}</p>
            </div>
            <div>
              <ThemeToggle />
            </div>
          </div>
        </div>
      </div>
    </IndexLayout>
  )
}
