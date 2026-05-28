import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { toast, ToastContainer } from "react-toastify"

import API from "../services/api"

const Login = () => {

  const navigate = useNavigate()

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleLogin = async (e) => {

    e.preventDefault()

    try {

      const response = await API.post(
        "/login",
        {
          email,
          password,
        }
      )

      console.log(response.data)

      // =========================
      // SAVE TOKEN
      // =========================

      localStorage.setItem(
        "token",
        response.data.token
      )

      localStorage.setItem(
        "role",
        response.data.role
      )

      // =========================
      // SUCCESS TOAST
      // =========================

      toast.success(
        "Login Successful"
      )

      // =========================
      // ROLE BASED NAVIGATION
      // =========================

      setTimeout(() => {

        if (
          response.data.role === "admin"
        ) {

          navigate("/admin")
        }

        else {

          navigate("/student")
        }

      }, 1500)

    } catch (error) {

      console.log(
        error.response?.data
      )

      const message =
        error.response?.data?.error ||
        "Invalid Email or Password"

      toast.error(message)
    }
  }

  return (

    <div
      className="flex items-center justify-center min-h-screen px-6 sm:px-0 bg-cover bg-center relative"
      style={{
        backgroundImage:
          "url('https://images.unsplash.com/photo-1522202176988-66273c2fd55f?q=80&w=2071&auto=format&fit=crop')",
      }}
    >

      {/* OVERLAY */}

      <div className="absolute inset-0 bg-black/50"></div>

      <ToastContainer />

      {/* LOGIN CARD */}

      <div className="relative z-10 bg-white/10 backdrop-blur-lg p-10 rounded-3xl shadow-2xl w-full sm:w-96 border border-white/20">

        {/* ========================= */}
        {/* TITLE */}
        {/* ========================= */}

        <h2 className="text-4xl font-bold text-white text-center mb-3">
          Login Account
        </h2>

        <p className="text-center text-gray-200 mb-8">
          Login to Student Management System
        </p>

        {/* ========================= */}
        {/* FORM */}
        {/* ========================= */}

        <form
          onSubmit={handleLogin}
          className="space-y-5"
        >

          {/* EMAIL */}

          <div>

            <label className="block text-sm font-medium text-white mb-2">
              Email
            </label>

            <input
              onChange={(e) =>
                setEmail(e.target.value)
              }
              value={email}
              className="w-full px-4 py-3 rounded-xl bg-white/20 border border-white/30 text-white placeholder-gray-300 outline-none focus:border-indigo-400"
              type="email"
              placeholder="Enter Email"
              required
            />

          </div>

          {/* PASSWORD */}

          <div>

            <label className="block text-sm font-medium text-white mb-2">
              Password
            </label>

            <input
              onChange={(e) =>
                setPassword(e.target.value)
              }
              value={password}
              className="w-full px-4 py-3 rounded-xl bg-white/20 border border-white/30 text-white placeholder-gray-300 outline-none focus:border-indigo-400"
              type="password"
              placeholder="Enter Password"
              required
            />

          </div>

          {/* FORGOT PASSWORD */}

          <div className="flex justify-end">

            <p className="text-sm text-indigo-200 cursor-pointer hover:underline">
              Forgot password?
            </p>

          </div>

          {/* BUTTON */}

          <button
            type="submit"
            className="w-full py-3 rounded-xl bg-gradient-to-r from-indigo-500 to-purple-600 text-white font-semibold hover:scale-105 transition duration-300"
          >
            Login
          </button>

        </form>

        {/* ========================= */}
        {/* FOOTER */}
        {/* ========================= */}

        <p className="text-center text-gray-200 mt-6">

          Don&apos;t have an account?

          <span
            onClick={() => navigate("/register")}
            className="text-white font-semibold cursor-pointer ml-1 hover:underline"
          >
            Register here
          </span>

        </p>

      </div>

    </div>
  )
}

export default Login