import {MouseEventHandler} from 'react'


interface ModalProps extends Readonly<{children: React.ReactNode}> {
  title?: string
  desc?: string
  closeModal: MouseEventHandler<any>
}

export default function Modal(props: ModalProps) {
  return (
    <section className="fixed w-screen h-screen bg-black/60 top-0 left-0">
      <div className="fixed w-[540px] max-h-[360px] bg-white left-1/2 -translate-x-1/2 top-40 rounded-xl border outline-neutral-500/20 outline py-4 px-14">
        <div className="flex justify-center mb-4">
          <div>
            <h3 className="text-lg font-bold text-center">{ props.title }</h3>
            <h5></h5>
          </div>
          <div className="absolute right-4 cursor-pointer" onClick={props.closeModal}><span className="material-symbols-rounded">close</span></div>
        </div>
        { props.children }
      </div>
    </section>
  )
}