import type { Comparison } from "@/types/fixture-analytis";
import React from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  LabelList,
} from "recharts";

interface ComparisonChartProps {
  comparison: Comparison;
  homeName: string;
  awayName: string;
}

const ComparisonChart: React.FC<ComparisonChartProps> = ({
  comparison,
  homeName,
  awayName,
}) => {
  const data = Object.keys(comparison).map((key) => ({
    name: key.toUpperCase().replace("_", " "),
    home: parseFloat(comparison[key as keyof Comparison].home),
    away: parseFloat(comparison[key as keyof Comparison].away),
  }));

  return (
    <div className="h-[400px] w-full bg-white p-6 rounded-3xl shadow-sm border border-slate-100">
      <div className="flex justify-between items-center mb-6">
        <h3 className="text-sm font-black text-slate-700 uppercase tracking-widest flex items-center gap-2">
          <svg
            className="w-5 h-5 text-indigo-500"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
            />
          </svg>
          Statistical Comparison (%)
        </h3>
      </div>
      <ResponsiveContainer width="100%" height="90%">
        <BarChart
          data={data}
          margin={{ top: 20, right: 10, left: -20, bottom: 20 }}
        >
          <CartesianGrid
            strokeDasharray="3 3"
            vertical={false}
            stroke="#f1f5f9"
          />
          <XAxis
            dataKey="name"
            axisLine={false}
            tickLine={false}
            tick={{ fill: "#64748b", fontSize: 10, fontWeight: 700 }}
          />
          <YAxis hide domain={[0, 115]} />
          <Tooltip
            cursor={{ fill: "#f8fafc" }}
            contentStyle={{
              borderRadius: "12px",
              border: "none",
              boxShadow: "0 10px 15px -3px rgb(0 0 0 / 0.1)",
            }}
          />
          <Legend
            iconType="circle"
            wrapperStyle={{
              paddingTop: "20px",
              fontSize: "12px",
              fontWeight: 600,
            }}
          />
          <Bar
            dataKey="home"
            name={homeName}
            fill="#6366f1"
            radius={[6, 6, 0, 0]}
          >
            <LabelList
              dataKey="home"
              position="top"
              style={{ fill: "#4338ca", fontSize: "11px", fontWeight: 800 }}
              formatter={(val: number) => `${val}%`}
            />
          </Bar>
          <Bar
            dataKey="away"
            name={awayName}
            fill="#ec4899"
            radius={[6, 6, 0, 0]}
          >
            <LabelList
              dataKey="away"
              position="top"
              style={{ fill: "#be185d", fontSize: "11px", fontWeight: 800 }}
              formatter={(val: number) => `${val}%`}
            />
          </Bar>
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default ComparisonChart;
