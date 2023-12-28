export function useViewTransition(func: any) {
  // 原始函数的参数args
  const _viewTransition = function () {
    document.startViewTransition(func)
  }
  return _viewTransition as any
}
