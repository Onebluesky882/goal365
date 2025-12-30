import type { FormData } from "@/types/predictions/types";

type FormChartProps = {
  formData: FormData;
  homeLabel?: string;
  awayLabel?: string;
};

const FormChart = ({
  formData,
  homeLabel = "เจ้าบ้าน",
  awayLabel = "ทีมเยือน",
}: FormChartProps) => {
  const chartData = [
    {
      label: "14 นัด",
      home: formData.home_form_14,
      away: formData.away_form_14,
      max: 42,
    },
    {
      label: "12 นัด",
      home: formData.home_form_12,
      away: formData.away_form_12,
      max: 36,
    },
    {
      label: "10 นัด",
      home: formData.home_form_10,
      away: formData.away_form_10,
      max: 30,
    },
    {
      label: "7 นัด",
      home: formData.home_form_7,
      away: formData.away_form_7,
      max: 21,
    },
    {
      label: "5 นัด",
      home: formData.home_form_5,
      away: formData.away_form_5,
      max: 15,
    },
  ];

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between text-sm">
        <div className="flex items-center gap-2">
          <div className="h-3 w-3 rounded-full bg-primary" />
          <span className="text-muted-foreground">{homeLabel}</span>
        </div>
        <div className="flex items-center gap-2">
          <span className="text-muted-foreground">{awayLabel}</span>
          <div className="h-3 w-3 rounded-full bg-secondary" />
        </div>
      </div>

      <div className="space-y-3">
        {chartData.map((item, index) => {
          const homePercent = (item.home / item.max) * 100;
          const awayPercent = (item.away / item.max) * 100;

          return (
            <div key={item.label} className="space-y-1">
              <div className="flex items-center justify-between text-xs text-muted-foreground">
                <span>{item.home} คะแนน</span>
                <span className="font-medium text-foreground">
                  {item.label}
                </span>
                <span>{item.away} คะแนน</span>
              </div>
              <div className="flex items-center gap-1 h-6">
                {/* Home bar - grows from right to left */}
                <div className="flex-1 flex justify-end">
                  <div
                    className="h-full rounded-l-md bg-primary origin-right transition-all duration-500"
                    style={{
                      width: `${homePercent}%`,
                      animationDelay: `${index * 100}ms`,
                    }}
                  />
                </div>

                <div className="w-px h-full bg-border" />

                {/* Away bar - grows from left to right */}
                <div className="flex-1">
                  <div
                    className="h-full rounded-r-md bg-secondary origin-left transition-all duration-500"
                    style={{
                      width: `${awayPercent}%`,
                      animationDelay: `${index * 100}ms`,
                    }}
                  />
                </div>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default FormChart;
