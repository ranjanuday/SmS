import React from "react"

const StatCard = ({
  title,
  value,
  icon,
  color,
}) => {

  return (

    <div className="bg-white/5 border border-white/10 backdrop-blur-lg rounded-3xl p-8 hover:scale-105 transition duration-300 shadow-xl">

      <div className="flex items-center justify-between">

        <div>

          <p className="text-gray-400 text-lg">
            {title}
          </p>

          <h2 className="text-5xl font-bold mt-4">
            {value}
          </h2>

        </div>

        <div
          className={`${color} p-4 rounded-2xl`}
        >
          {icon}
        </div>

      </div>

    </div>
  )
}

export default StatCard