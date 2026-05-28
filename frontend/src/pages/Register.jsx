import { useState } from "react"
import { useNavigate } from "react-router-dom"
import { toast, ToastContainer } from "react-toastify"

import API from "../services/api"

function Register() {

  const navigate = useNavigate()

  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [role, setRole] = useState("student")

  const handleRegister = async (e) => {

    e.preventDefault()

    try {

      const response = await API.post(
        "/register",
        {
          name,
          email,
          password,
          role,
        }
      )

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
        "Registration Successful"
      )

      // =========================
      // NAVIGATE TO LOGIN
      // =========================

      setTimeout(() => {

        navigate("/")

      }, 1500)

    } catch (error) {

      console.log(
        error.response?.data
      )

      const message =
        error.response?.data?.error ||
        "Registration Failed"

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

      {/* REGISTER CARD */}

      <div className="relative z-10 bg-white/10 backdrop-blur-lg p-10 rounded-3xl shadow-2xl w-full sm:w-[420px] border border-white/20">

        {/* TITLE */}

        <h1 className="text-4xl font-bold text-center text-white mb-2">
          Create Account
        </h1>

        <p className="text-center text-gray-200 mb-8">
          Register to Student Management System
        </p>

        {/* FORM */}

        <form
          onSubmit={handleRegister}
          className="space-y-5"
        >

          {/* NAME */}

          <div>

            <label className="block mb-2 font-medium text-white">
              Name
            </label>

            <input
              type="text"
              placeholder="Enter your name"
              className="w-full px-4 py-3 rounded-xl bg-white/20 border border-white/30 text-white placeholder-gray-300 outline-none "
              value={name}
              onChange={(e) =>
                setName(e.target.value)
              }
              required
            />

          </div>

          {/* EMAIL */}

          <div>

            <label className="block mb-2 font-medium text-white">
              Email
            </label>

            <input
              type="email"
              placeholder="Enter your email"
              className="w-full px-4 py-3 rounded-xl bg-white/20 border border-white/30 text-white placeholder-gray-300 outline-none "
              value={email}
              onChange={(e) =>
                setEmail(e.target.value)
              }
              required
            />

          </div>

          {/* PASSWORD */}

          <div>

            <label className="block mb-2 font-medium text-white">
              Password
            </label>

            <input
              type="password"
              placeholder="Enter your password"
              className="w-full px-4 py-3 rounded-xl bg-white/20 border border-white/30 text-white placeholder-gray-300 outline-none "
              value={password}
              onChange={(e) =>
                setPassword(e.target.value)
              }
              required
            />

          </div>

          {/* ROLE */}

          <div>

            <label className="block mb-2 font-medium text-white">
              Role
            </label>

            <select
              className="w-full px-4 py-3 rounded-xl bg-white/20 border border-white/30 text-white outline-none"
              value={role}
              onChange={(e) =>
                setRole(e.target.value)
              }
            >

              <option
                value="student"
                className="text-black"
              >
                Student
              </option>

              <option
                value="admin"
                className="text-black"
              >
                Admin
              </option>

            </select>

          </div>

          {/* BUTTON */}

          <button
            type="submit"
             className="w-full py-3 rounded-xl bg-gradient-to-r from-indigo-500 to-purple-600 text-white font-semibold hover:scale-105 transition duration-300"
          >
            Register
          </button>

        </form>

        {/* FOOTER */}

        <p className="text-center text-gray-200 mt-6">

          Already have an account?

          <span
            className="text-white font-semibold cursor-pointer ml-1 hover:underline"
            onClick={() => navigate("/")}
          >
            Login
          </span>

        </p>

      </div>

    </div>
  )
}

export default Register