import { type ReactNode, useState } from 'react'
import { type Point } from './point'

export interface SelectableProps {
    children: ReactNode
    onSelected?: (selected: boolean) => void
}

export function Selectable(props: SelectableProps) {
    const [selected, setSelected] = useState<boolean>(false)
    const [loc, setLoc] = useState<Point>({ x: 0, y: 0 })
    const [bBox, setBBox] = useState<DOMRect>()

    const onClick = () => {
        setSelected(!selected)

        if (props.onSelected) {
            props.onSelected(!selected)
        }
    }

    return <>
        <g ref={el => {
            if (!el || bBox) {
                return
            }

            setBBox(el.getBBox())
        }} onMouseDown={(e) => {
            setLoc({ x: e.clientX, y: e.clientY })
        }}
        onMouseUp={(e) => {
            if (loc.x === e.clientX && loc.y === e.clientY) {
                onClick()
            }
        }}>
            {!selected && props.children}
            {selected && <g>
                {bBox && <rect x={-3}
                    y={-3}
                    width={bBox?.width + 8}
                    height={bBox?.height + 9}
                    strokeWidth="3"
                    stroke="rgb(20, 18, 230)"
                    fill="#fff"
                    strokeDasharray="5,2,2,2,2,2"
                />}
                {props.children}
            </g>}
        </g>
    </>
}
